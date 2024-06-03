package hotel

import (
	"github.com/bastean/bookingo/pkg/context/hotel/application/login"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type LoginQuery = login.Query

func NewLogin(repository model.Repository, hashing models.Hashing) *login.QueryHandler {
	useCase := &login.Login{
		Repository: repository,
		Hashing:    hashing,
	}

	return &login.QueryHandler{
		UseCase: useCase,
	}
}
