package services

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

func IsPasswordInvalid(hashing models.Hashing, hashed, plain string) error {
	if hashing.IsNotEqual(hashed, plain) {
		return errors.NewFailure(&errors.Bubble{
			Where: "IsPasswordInvalid",
			What:  "passwords do not match",
		})
	}

	return nil
}
