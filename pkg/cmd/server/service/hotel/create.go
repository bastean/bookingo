package hotel

import (
	"github.com/bastean/bookingo/pkg/context/hotel/application/create"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type CreateCommand = create.Command

func NewCreate(repository model.Repository, broker models.Broker) *create.CommandHandler {
	useCase := &create.Create{
		Repository: repository,
	}

	return &create.CommandHandler{
		UseCase: useCase,
		Broker:  broker,
	}
}
