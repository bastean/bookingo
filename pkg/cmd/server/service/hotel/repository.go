package hotel

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/hotel/infrastructure/persistence"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/persistences"
)

func NewMongoCollection(database *persistences.MongoDB, name string, hashing models.Hashing) (model.Repository, error) {
	collection, err := persistence.NewMongoCollection(database, name, hashing)

	if err != nil {
		return nil, errors.BubbleUp(err, "NewMongoCollection")
	}

	return collection, nil
}
