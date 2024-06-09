package persistence

import (
	"context"

	"github.com/bastean/bookingo/pkg/context/booking/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/booking/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/persistences"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BookingDocument struct {
	HotelID   string  `bson:"hotelId,omitempty"`
	ID        string  `bson:"id,omitempty"`
	FirstName string  `bson:"firstName,omitempty"`
	LastName  string  `bson:"lastName,omitempty"`
	Email     string  `bson:"email,omitempty"`
	Phone     string  `bson:"phone,omitempty"`
	CheckIn   string  `bson:"checkIn,omitempty"`
	CheckOut  string  `bson:"checkOut,omitempty"`
	Room      string  `bson:"room,omitempty"`
	Currency  string  `bson:"currency,omitempty"`
	Total     float32 `bson:"total,omitempty"`
}

type BookingCollection struct {
	collection *mongo.Collection
}

func (db *BookingCollection) Save(booking *aggregate.Booking) error {
	newBooking := BookingDocument(*booking.ToPrimitives())

	_, err := db.collection.InsertOne(context.Background(), newBooking)

	if mongo.IsDuplicateKeyError(err) {
		return errors.BubbleUp(persistences.HandleMongoDuplicateKeyError(err), "Save")
	}

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Save",
			What:  "failure to save a booking",
			Why: errors.Meta{
				"HotelID": booking.HotelID.Value(),
				"ID":      booking.ID.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *BookingCollection) Update(booking *aggregate.Booking) error {
	updatedBooking := BookingDocument(*booking.ToPrimitives())

	filter := bson.D{{Key: "id", Value: booking.ID.Value()}}

	_, err := db.collection.ReplaceOne(context.Background(), filter, updatedBooking)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Update",
			What:  "failure to update a booking",
			Why: errors.Meta{
				"HotelID": booking.HotelID.Value(),
				"ID":      booking.ID.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *BookingCollection) Delete(id models.ValueObject[string]) error {
	filter := bson.D{{Key: "id", Value: id.Value()}}

	_, err := db.collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Delete",
			What:  "failure to delete a booking",
			Why: errors.Meta{
				"ID": id.Value(),
			},
			Who: err,
		})
	}

	return nil
}

func (db *BookingCollection) Search(criteria *model.RepositorySearchCriteria) (*aggregate.Booking, error) {
	var filter bson.D
	var index string

	switch {
	case criteria.HotelID != nil:
		filter = bson.D{{Key: "hotel_id", Value: criteria.HotelID.Value()}}
		index = criteria.ID.Value()
	case criteria.ID != nil:
		filter = bson.D{{Key: "id", Value: criteria.ID.Value()}}
		index = criteria.ID.Value()
	case criteria.Email != nil:
		filter = bson.D{{Key: "email", Value: criteria.Email.Value()}}
		index = criteria.Email.Value()
	case criteria.Phone != nil:
		filter = bson.D{{Key: "phone", Value: criteria.Phone.Value()}}
		index = criteria.Phone.Value()
	}

	result := db.collection.FindOne(context.Background(), filter)

	if err := result.Err(); err != nil {
		return nil, persistences.HandleMongoDocumentNotFound(index, err)
	}

	bookingPrimitive := new(aggregate.BookingPrimitive)

	result.Decode(bookingPrimitive)

	booking, err := aggregate.FromPrimitives(bookingPrimitive)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Search",
			What:  "failure to search for a booking",
			Why: errors.Meta{
				"Index": index,
			},
			Who: err,
		})
	}

	return booking, nil
}

func NewMongoCollection(mdb *persistences.MongoDB, collectionName string) (model.Repository, error) {
	collection := mdb.Database.Collection(collectionName)

	_, err := collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "NewMongoCollection",
			What:  "failure to create indexes for booking collection",
			Why: errors.Meta{
				"Collection": collectionName,
			},
			Who: err,
		})
	}

	return &BookingCollection{
		collection: collection,
	}, nil
}
