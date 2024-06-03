package hotel

import (
	"github.com/bastean/bookingo/pkg/context/hotel/application/update"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type UpdateCommand = update.Command

func NewUpdate(repository model.Repository, hashing models.Hashing) *update.CommandHandler {
	useCase := &update.Update{
		Repository: repository,
		Hashing:    hashing,
	}

	return &update.CommandHandler{
		UseCase: useCase,
	}
}
