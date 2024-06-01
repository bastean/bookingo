package user

import (
	"github.com/bastean/bookingo/pkg/context/user/application/read"
	"github.com/bastean/bookingo/pkg/context/user/domain/model"
)

type ReadQuery = read.Query

type ReadResponse = read.Response

func NewRead(repository model.Repository) *read.QueryHandler {
	useCase := &read.Read{
		Repository: repository,
	}

	return &read.QueryHandler{
		UseCase: useCase,
	}
}
