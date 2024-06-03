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

type hotelDocument struct {
	Id        string `bson:"id"`
	Email     string `bson:"email"`
	Hotelname string `bson:"hotelname"`
	Password  string `bson:"password"`
	Verified  bool   `bson:"verified"`
}

type HotelCollection struct {
	collection *mongo.Collection
	hashing    models.Hashing
}

func (db *HotelCollection) Save(hotel *aggregate.Hotel) error {
	newHotel := hotelDocument(*hotel.ToPrimitives())

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
				"Id": hotel.Id.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *HotelCollection) Update(hotel *aggregate.Hotel) error {
	updateFilter := bson.M{"id": hotel.Id.Value()}

	updateHotel := bson.M{}

	if hotel.Email != nil {
		updateHotel["email"] = hotel.Email.Value()
	}

	if hotel.Hotelname != nil {
		updateHotel["hotelname"] = hotel.Hotelname.Value()
	}

	if hotel.Password != nil {
		hashed, err := db.hashing.Hash(hotel.Password.Value())

		if err != nil {
			return errors.BubbleUp(err, "Update")
		}

		updateHotel["password"] = hashed
	}

	if hotel.Verified != nil {
		updateHotel["verified"] = hotel.Verified.Value()
	}

	_, err := db.collection.UpdateOne(context.Background(), updateFilter, bson.M{"$set": updateHotel})

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Update",
			What:  "failure to update a hotel",
			Why: errors.Meta{
				"Id": hotel.Id.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *HotelCollection) Delete(id models.ValueObject[string]) error {
	deleteFilter := bson.M{"id": id.Value()}

	_, err := db.collection.DeleteOne(context.Background(), deleteFilter)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Delete",
			What:  "failure to delete a hotel",
			Why: errors.Meta{
				"Id": id.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *HotelCollection) Search(filter model.RepositorySearchCriteria) (*aggregate.Hotel, error) {
	var searchFilter bson.M
	var index string

	if filter.Email != nil {
		searchFilter = bson.M{"email": filter.Email.Value()}
		index = filter.Email.Value()
	}

	if filter.Id != nil {
		searchFilter = bson.M{"id": filter.Id.Value()}
		index = filter.Id.Value()
	}

	result := db.collection.FindOne(context.Background(), searchFilter)

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
				"Id":    filter.Id.Value(),
				"Email": filter.Email.Value(),
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
			Keys:    bson.D{{Key: "hotelname", Value: 1}},
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
