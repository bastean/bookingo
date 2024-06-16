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

func create(primitive *HotelPrimitive) (*Hotel, error) {
	aggregateRoot := aggregates.NewAggregateRoot()

	idVO, errId := valueobj.NewId(primitive.Id)
	nameVO, errName := valueobj.NewName(primitive.Name)
	emailVO, errEmail := valueobj.NewEmail(primitive.Email)
	phoneVO, errPhone := valueobj.NewPhone(primitive.Phone)
	passwordVO, errPassword := valueobj.NewPassword(primitive.Password)
	verifiedVO, errVerified := valueobj.NewVerified(primitive.Verified)

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
	hotel, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromPrimitives")
	}

	return hotel, nil
}

func NewHotel(primitive *HotelPrimitive) (*Hotel, error) {
	primitive.Verified = false

	hotel, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "NewHotel")
	}

	message, err := event.NewCreatedSucceeded(&event.CreatedSucceeded{
		Attributes: &event.CreatedSucceededAttributes{
			Id:    primitive.Id,
			Name:  primitive.Name,
			Email: primitive.Email,
			Phone: primitive.Phone,
		},
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "NewHotel")
	}

	hotel.RecordMessage(message)

	return hotel, nil
}
