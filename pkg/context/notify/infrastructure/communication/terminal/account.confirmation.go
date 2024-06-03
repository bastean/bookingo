package terminal

import (
	"fmt"

	"github.com/bastean/bookingo/pkg/context/notify/application/send"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

type AccountConfirmation struct {
	models.Logger
	ServerURL string
}

func (client *AccountConfirmation) Submit(data any) error {
	account := data.(*send.CreatedSucceededEventAttributes)

	confirmationLink := fmt.Sprintf("Hi %s, please confirm your account through this link: %s/verify/%s", account.Hotelname, client.ServerURL, account.Id)

	client.Logger.Info(confirmationLink)

	return nil
}
