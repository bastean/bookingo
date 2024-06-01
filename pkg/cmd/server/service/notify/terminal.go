package notify

import (
	"github.com/bastean/bookingo/pkg/context/notify/domain/model"
	"github.com/bastean/bookingo/pkg/context/notify/infrastructure/communication/terminal"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

func NewTerminalAccountConfirmation(logger models.Logger, serverURL string) model.Transport {
	return &terminal.AccountConfirmation{
		Logger:    logger,
		ServerURL: serverURL,
	}
}
