package persistence

import (
	"context"

	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/persistences"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HotelDocument struct {
	ID       string `bson:"id,omitempty"`
	Name     string `bson:"name,omitempty"`
	Email    string `bson:"email,omitempty"`
	Phone    string `bson:"phone,omitempty"`
	Password string `bson:"password,omitempty"`
	Verified bool   `bson:"verified,omitempty"`
}

type HotelCollection struct {
	collection *mongo.Collection
	hashing    models.Hashing
}

func (db *HotelCollection) Save(hotel *aggregate.Hotel) error {
	newHotel := HotelDocument(*hotel.ToPrimitives())

	hashed, err := db.hashing.Hash(newHotel.Password)

	if err != nil {
		return errors.BubbleUp(err, "Save")
	}

	newHotel.Password = hashed

	_, err = db.collection.InsertOne(context.Background(), newHotel)

	if mongo.IsDuplicateKeyError(err) {
		return errors.BubbleUp(persistences.HandleMongoDuplicateKeyError(err), "Save")
	}

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Save",
			What:  "failure to save a hotel",
			Why: errors.Meta{
				"ID": hotel.ID.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *HotelCollection) Update(hotel *aggregate.Hotel) error {
	updatedHotel := HotelDocument(*hotel.ToPrimitives())

	filter := bson.D{{Key: "id", Value: hotel.ID.Value()}}

	hashed, err := db.hashing.Hash(hotel.Password.Value())

	if err != nil {
		return errors.BubbleUp(err, "Update")
	}

	updatedHotel.Password = hashed

	_, err = db.collection.ReplaceOne(context.Background(), filter, updatedHotel)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Update",
			What:  "failure to update a hotel",
			Why: errors.Meta{
				"ID": hotel.ID.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *HotelCollection) Verify(id models.ValueObject[string]) error {
	filter := bson.D{{Key: "id", Value: id.Value()}}

	_, err := db.collection.UpdateOne(context.Background(), filter, bson.D{{Key: "$set", Value: bson.D{
		{Key: "verified", Value: true},
	}}})

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Verify",
			What:  "failure to verify a hotel",
			Why: errors.Meta{
				"ID": id.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *HotelCollection) Delete(id models.ValueObject[string]) error {
	filter := bson.D{{Key: "id", Value: id.Value()}}

	_, err := db.collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Delete",
			What:  "failure to delete a hotel",
			Why: errors.Meta{
				"ID": id.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *HotelCollection) Search(criteria model.RepositorySearchCriteria) (*aggregate.Hotel, error) {
	var filter bson.D
	var index string

	if criteria.Phone != nil {
		filter = bson.D{{Key: "phone", Value: criteria.Phone.Value()}}
		index = criteria.Phone.Value()
	}

	if criteria.Email != nil {
		filter = bson.D{{Key: "email", Value: criteria.Email.Value()}}
		index = criteria.Email.Value()
	}

	if criteria.ID != nil {
		filter = bson.D{{Key: "id", Value: criteria.ID.Value()}}
		index = criteria.ID.Value()
	}

	result := db.collection.FindOne(context.Background(), filter)

	if err := result.Err(); err != nil {
		return nil, persistences.HandleMongoDocumentNotFound(index, err)
	}

	hotelPrimitive := new(aggregate.HotelPrimitive)

	result.Decode(hotelPrimitive)

	hotel, err := aggregate.FromPrimitives(hotelPrimitive)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Search",
			What:  "failure to search for a hotel",
			Why: errors.Meta{
				"Index": index,
			},
			Who: err,
		})
	}

	return hotel, nil
}

func NewMongoCollection(mdb *persistences.MongoDB, collectionName string, hashing models.Hashing) (model.Repository, error) {
	collection := mdb.Database.Collection(collectionName)

	_, err := collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "phone", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "NewMongoCollection",
			What:  "failure to create indexes for hotel collection",
			Why: errors.Meta{
				"Collection": collectionName,
			},
			Who: err,
		})
	}

	return &HotelCollection{
		collection: collection,
		hashing:    hashing,
	}, nil
}
