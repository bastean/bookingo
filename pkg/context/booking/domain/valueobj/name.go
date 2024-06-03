package valueobj

import (
	"regexp"
	"strings"

	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

const NameMinCharactersLength = "1"
const NameMaxCharactersLength = "64"

type Name struct {
	Name string
}

func (name *Name) Value() string {
	return name.Name
}

func (name *Name) IsValid() error {
	validate := regexp.MustCompile(`^[A-Za-z]{1,64}$`)

	if !validate.MatchString(name.Name) {
		return errors.Default()
	}

	return nil
}

func NewName(name string) (models.ValueObject[string], error) {
	name = strings.TrimSpace(name)

	nameVO := &Name{
		Name: name,
	}

	if nameVO.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewName",
			What:  "name must be between " + NameMinCharactersLength + " to " + NameMaxCharactersLength + " characters",
			Why: errors.Meta{
				"Name": name,
			},
		})
	}

	return nameVO, nil
}
