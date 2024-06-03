package aggregate

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/message"
	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/aggregates"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type Booking struct {
	*aggregates.AggregateRoot
	HotelID   models.ValueObject[string]
	ID        models.ValueObject[string]
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
	HotelID   string
	ID        string
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

func create(hotelID, id, firstName, lastName, email, phone, checkIn, checkOut, room, currency string, total float32) (*Booking, error) {
	aggregateRoot := aggregates.NewAggregateRoot()

	hotelIdVO, errHotelID := valueobj.NewId(hotelID)
	idVO, errID := valueobj.NewId(id)
	firstNameVO, errFirstName := valueobj.NewName(firstName)
	lastNameVO, errLastName := valueobj.NewName(lastName)
	emailVO, errEmail := valueobj.NewEmail(email)
	phoneVO, errPhone := valueobj.NewPhone(phone)
	checkInVO, errCheckIn := valueobj.NewCheck(checkIn)
	checkOutVO, errCheckOut := valueobj.NewCheck(checkOut)
	roomVO, errRoom := valueobj.NewRoom(room)
	currencyVO, errCurrency := valueobj.NewCurrency(currency)
	totalVO, errTotal := valueobj.NewTotal(total)

	err := errors.Join(errHotelID, errID, errFirstName, errLastName, errEmail, errPhone, errCheckIn, errCheckOut, errRoom, errCurrency, errTotal)

	if err != nil {
		return nil, errors.BubbleUp(err, "create")
	}

	return &Booking{
		AggregateRoot: aggregateRoot,
		HotelID:       hotelIdVO,
		ID:            idVO,
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
		HotelID:   booking.HotelID.Value(),
		ID:        booking.ID.Value(),
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
	booking, err := create(
		primitive.HotelID,
		primitive.ID,
		primitive.FirstName,
		primitive.LastName,
		primitive.Email,
		primitive.Phone,
		primitive.CheckIn,
		primitive.CheckOut,
		primitive.Room,
		primitive.Currency,
		primitive.Total,
	)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromPrimitives")
	}

	return booking, nil
}

func NewBooking(hotelID, id, firstName, lastName, email, phone, checkIn, checkOut, room, currency string, total float32) (*Booking, error) {
	booking, err := create(
		hotelID,
		id,
		firstName,
		lastName,
		email,
		phone,
		checkIn,
		checkOut,
		room,
		currency,
		total,
	)

	if err != nil {
		return nil, errors.BubbleUp(err, "NewBooking")
	}

	eventMessage, err := message.NewCreatedSucceededEvent(&message.CreatedSucceededEventAttributes{
		HotelID:   hotelID,
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		CheckIn:   checkIn,
		CheckOut:  checkOut,
		Room:      room,
		Currency:  currency,
		Total:     total,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "NewBooking")
	}

	booking.RecordMessage(eventMessage)

	return booking, nil
}
