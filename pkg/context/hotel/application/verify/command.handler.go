package verify

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type CommandHandler struct {
	models.UseCase[models.ValueObject[string], *types.Empty]
}

func (handler *CommandHandler) Handle(command *Command) error {
	idVO, err := valueobj.NewId(command.Id)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	_, err = handler.UseCase.Run(idVO)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}
