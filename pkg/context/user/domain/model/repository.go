package model

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/user/domain/aggregate"
)

type RepositorySearchCriteria struct {
	Id    models.ValueObject[string]
	Email models.ValueObject[string]
}

type Repository interface {
	Save(user *aggregate.User) error
	Update(user *aggregate.User) error
	Delete(id models.ValueObject[string]) error
	Search(filter RepositorySearchCriteria) (*aggregate.User, error)
}
