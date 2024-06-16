package aggregate

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/event"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/aggregates"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type Hotel struct {
	*aggregates.AggregateRoot
	Id       models.ValueObject[string]
	Name     models.ValueObject[string]
	Email    models.ValueObject[string]
	Phone    models.ValueObject[string]
	Password models.ValueObject[string]
	Verified models.ValueObject[bool]
}

type HotelPrimitive struct {
	Id       string
	Name     string
	Email    string
	Phone    string
	Password string
	Verified bool
}

func create(id, name, email, phone, password string, verified bool) (*Hotel, error) {
	aggregateRoot := aggregates.NewAggregateRoot()

	idVO, errId := valueobj.NewId(id)
	nameVO, errName := valueobj.NewName(name)
	emailVO, errEmail := valueobj.NewEmail(email)
	phoneVO, errPhone := valueobj.NewPhone(phone)
	passwordVO, errPassword := valueobj.NewPassword(password)
	verifiedVO, errVerified := valueobj.NewVerified(verified)

	err := errors.Join(errId, errName, errEmail, errPhone, errPassword, errVerified)

	if err != nil {
		return nil, errors.BubbleUp(err, "create")
	}

	return &Hotel{
		AggregateRoot: aggregateRoot,
		Id:            idVO,
		Name:          nameVO,
		Email:         emailVO,
		Phone:         phoneVO,
		Password:      passwordVO,
		Verified:      verifiedVO,
	}, nil
}

func (hotel *Hotel) ToPrimitives() *HotelPrimitive {
	return &HotelPrimitive{
		Id:       hotel.Id.Value(),
		Name:     hotel.Name.Value(),
		Email:    hotel.Email.Value(),
		Phone:    hotel.Phone.Value(),
		Password: hotel.Password.Value(),
		Verified: hotel.Verified.Value(),
	}
}

func FromPrimitives(primitive *HotelPrimitive) (*Hotel, error) {
	hotel, err := create(
		primitive.Id,
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

	message, err := event.NewCreatedSucceeded(&event.CreatedSucceeded{
		Attributes: &event.CreatedSucceededAttributes{
			Id:    id,
			Name:  name,
			Email: email,
			Phone: phone,
		},
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "NewHotel")
	}

	hotel.RecordMessage(message)

	return hotel, nil
}
