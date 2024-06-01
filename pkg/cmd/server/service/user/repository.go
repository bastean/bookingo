package user

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/persistences"
	"github.com/bastean/bookingo/pkg/context/user/domain/model"
	"github.com/bastean/bookingo/pkg/context/user/infrastructure/persistence"
)

func NewMongoCollection(database *persistences.MongoDB, name string, hashing model.Hashing) (model.Repository, error) {
	collection, err := persistence.NewMongoCollection(database, name, hashing)

	if err != nil {
		return nil, errors.BubbleUp(err, "NewMongoCollection")
	}

	return collection, nil
}
