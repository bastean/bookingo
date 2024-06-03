package hotel

import (
	"github.com/bastean/bookingo/pkg/context/hotel/application/delete"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type DeleteCommand = delete.Command

func NewDelete(repository model.Repository, hashing models.Hashing) *delete.CommandHandler {
	useCase := &delete.Delete{
		Repository: repository,
		Hashing:    hashing,
	}

	return &delete.CommandHandler{
		UseCase: useCase,
	}
}
