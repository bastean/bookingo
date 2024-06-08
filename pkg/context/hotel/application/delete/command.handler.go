package delete

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
)

type Input struct {
	ID, Password models.ValueObject[string]
}

type CommandHandler struct {
	models.UseCase[*Input, types.Empty]
}

func (handler *CommandHandler) Handle(command *Command) error {
	id, errId := valueobj.NewId(command.ID)
	password, errPassword := valueobj.NewPassword(command.Password)

	err := errors.Join(errId, errPassword)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	_, err = handler.UseCase.Run(&Input{
		ID:       id,
		Password: password,
	})

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}
