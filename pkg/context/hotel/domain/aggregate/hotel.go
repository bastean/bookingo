package aggregate

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/message"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/aggregates"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type Hotel struct {
	*aggregates.AggregateRoot
	ID       models.ValueObject[string]
	Name     models.ValueObject[string]
	Email    models.ValueObject[string]
	Phone    models.ValueObject[string]
	Password models.ValueObject[string]
	Verified models.ValueObject[bool]
}

type HotelPrimitive struct {
	ID       string
	Name     string
	Email    string
	Phone    string
	Password string
	Verified bool
}

func create(id, name, email, phone, password string, verified bool) (*Hotel, error) {
	aggregateRoot := aggregates.NewAggregateRoot()

	idVO, errID := valueobj.NewId(id)
	nameVO, errName := valueobj.NewName(name)
	emailVO, errEmail := valueobj.NewEmail(email)
	phoneVO, errPhone := valueobj.NewPhone(phone)
	passwordVO, errPassword := valueobj.NewPassword(password)
	verifiedVO, errVerified := valueobj.NewVerified(verified)

	err := errors.Join(errID, errName, errEmail, errPhone, errPassword, errVerified)

	if err != nil {
		return nil, errors.BubbleUp(err, "create")
	}

	return &Hotel{
		AggregateRoot: aggregateRoot,
		ID:            idVO,
		Name:          nameVO,
		Email:         emailVO,
		Phone:         phoneVO,
		Password:      passwordVO,
		Verified:      verifiedVO,
	}, nil
}

func (hotel *Hotel) ToPrimitives() *HotelPrimitive {
	return &HotelPrimitive{
		ID:       hotel.ID.Value(),
		Name:     hotel.Name.Value(),
		Email:    hotel.Email.Value(),
		Phone:    hotel.Phone.Value(),
		Password: hotel.Password.Value(),
		Verified: hotel.Verified.Value(),
	}
}

func FromPrimitives(primitive *HotelPrimitive) (*Hotel, error) {
	hotel, err := create(
		primitive.ID,
		primitive.Name,
		primitive.Email,
		primitive.Phone,
		primitive.Password,
		primitive.Verified,
	)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromPrimitives")
	}

	return hotel, nil
}

func NewHotel(id, name, email, phone, password string) (*Hotel, error) {
	verified := false

	hotel, err := create(
		id,
		name,
		email,
		phone,
		password,
		verified,
	)

	if err != nil {
		return nil, errors.BubbleUp(err, "NewHotel")
	}

	eventMessage, err := message.NewCreatedSucceededEvent(&message.CreatedSucceededEventAttributes{
		ID:    id,
		Name:  name,
		Email: email,
		Phone: phone,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "NewHotel")
	}

	hotel.RecordMessage(eventMessage)

	return hotel, nil
}
