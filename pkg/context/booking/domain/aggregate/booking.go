package aggregate

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/event"
	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/aggregates"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type Booking struct {
	*aggregates.AggregateRoot
	HotelId   models.ValueObject[string]
	Id        models.ValueObject[string]
	FirstName models.ValueObject[string]
	LastName  models.ValueObject[string]
	Email     models.ValueObject[string]
	Phone     models.ValueObject[string]
	CheckIn   models.ValueObject[string]
	CheckOut  models.ValueObject[string]
	Room      models.ValueObject[string]
	Currency  models.ValueObject[string]
	Total     models.ValueObject[float32]
}

type BookingPrimitive struct {
	HotelId   string
	Id        string
	FirstName string
	LastName  string
	Email     string
	Phone     string
	CheckIn   string
	CheckOut  string
	Room      string
	Currency  string
	Total     float32
}

func create(primitive *BookingPrimitive) (*Booking, error) {
	aggregateRoot := aggregates.NewAggregateRoot()

	hotelIdVO, errHotelId := valueobj.NewId(primitive.HotelId)
	idVO, errId := valueobj.NewId(primitive.Id)
	firstNameVO, errFirstName := valueobj.NewName(primitive.FirstName)
	lastNameVO, errLastName := valueobj.NewName(primitive.LastName)
	emailVO, errEmail := valueobj.NewEmail(primitive.Email)
	phoneVO, errPhone := valueobj.NewPhone(primitive.Phone)
	checkInVO, errCheckIn := valueobj.NewCheck(primitive.CheckIn)
	checkOutVO, errCheckOut := valueobj.NewCheck(primitive.CheckOut)
	roomVO, errRoom := valueobj.NewRoom(primitive.Room)
	currencyVO, errCurrency := valueobj.NewCurrency(primitive.Currency)
	totalVO, errTotal := valueobj.NewTotal(primitive.Total)

	err := errors.Join(errHotelId, errId, errFirstName, errLastName, errEmail, errPhone, errCheckIn, errCheckOut, errRoom, errCurrency, errTotal)

	if err != nil {
		return nil, errors.BubbleUp(err, "create")
	}

	return &Booking{
		AggregateRoot: aggregateRoot,
		HotelId:       hotelIdVO,
		Id:            idVO,
		FirstName:     firstNameVO,
		LastName:      lastNameVO,
		Email:         emailVO,
		Phone:         phoneVO,
		CheckIn:       checkInVO,
		CheckOut:      checkOutVO,
		Room:          roomVO,
		Currency:      currencyVO,
		Total:         totalVO,
	}, nil
}

func (booking *Booking) ToPrimitives() *BookingPrimitive {
	return &BookingPrimitive{
		HotelId:   booking.HotelId.Value(),
		Id:        booking.Id.Value(),
		FirstName: booking.FirstName.Value(),
		LastName:  booking.LastName.Value(),
		Email:     booking.Email.Value(),
		Phone:     booking.Phone.Value(),
		CheckIn:   booking.CheckIn.Value(),
		CheckOut:  booking.CheckOut.Value(),
		Room:      booking.Room.Value(),
		Currency:  booking.Currency.Value(),
		Total:     booking.Total.Value(),
	}
}

func FromPrimitives(primitive *BookingPrimitive) (*Booking, error) {
	booking, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromPrimitives")
	}

	return booking, nil
}

func NewBooking(primitive *BookingPrimitive) (*Booking, error) {
	booking, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "NewBooking")
	}

	attributes := event.CreatedSucceededAttributes(*primitive)

	message, err := event.NewCreatedSucceeded(&event.CreatedSucceeded{
		Attributes: &attributes,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "NewBooking")
	}

	booking.RecordMessage(message)

	return booking, nil
}
