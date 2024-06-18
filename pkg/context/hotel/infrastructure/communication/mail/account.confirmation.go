package mail

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"

	"github.com/bastean/bookingo/pkg/context/hotel/domain/event"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/transports"
)

type AccountConfirmation struct {
	*transports.SMTP
}

func (client *AccountConfirmation) Submit(data any) error {
	hotel, ok := data.(*event.CreatedSucceededAttributes)

	if !ok {
		return errors.NewInternal(&errors.Bubble{
			Where: "Submit",
			What:  "failure in type assertion",
			Why: errors.Meta{
				"Expected": new(event.CreatedSucceededAttributes),
				"Actual":   data,
			},
		})
	}

	var message bytes.Buffer

	headers := fmt.Sprintf("From: %s\n"+"To: %s\n"+"Subject: Account Confirmation", client.Username, hotel.Email)

	message.Write([]byte(fmt.Sprintf("%s\n%s\n", headers, client.MIMEHeaders)))

	confirmationLink := fmt.Sprintf("%s/verify/%s", client.ServerURL, hotel.Id)

	AccountConfirmationTemplate(hotel.Name, confirmationLink).Render(context.Background(), &message)

	err := smtp.SendMail(client.SMTPServerURL, client.Auth, client.Username, []string{hotel.Email}, message.Bytes())

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Submit",
			What:  "failure to send an account confirmation mail",
			Why: errors.Meta{
				"Hotel Id":        hotel.Id,
				"SMTP Server URL": client.SMTPServerURL,
			},
			Who: err,
		})
	}

	return nil
}
