package login

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/aggregate"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type Input struct {
	Email    models.ValueObject[string]
	Password models.ValueObject[string]
}

type QueryHandler struct {
	models.UseCase[*Input, *aggregate.Hotel]
}

func (handler *QueryHandler) Handle(query *Query) (*Response, error) {
	email, errEmail := valueobj.NewEmail(query.Email)
	password, errPassword := valueobj.NewPassword(query.Password)

	err := errors.Join(errEmail, errPassword)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	hotel, err := handler.UseCase.Run(&Input{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	response := Response(*hotel.ToPrimitives())

	return &response, nil
}
