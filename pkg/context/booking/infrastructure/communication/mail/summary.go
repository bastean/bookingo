package mail

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"

	"github.com/bastean/bookingo/pkg/context/booking/domain/event"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/transports"
)

type Summary struct {
	*transports.SMTP
}

func (client *Summary) Submit(data any) error {
	booking, ok := data.(*event.CreatedSucceededAttributes)

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

	headers := fmt.Sprintf("From: %v\n"+"To: %v\n"+"Subject: Booking Summary", client.Username, booking.Email)

	message.Write([]byte(fmt.Sprintf("%v\n%v\n", headers, client.MIMEHeaders)))

	SummaryTemplate(booking).Render(context.Background(), &message)

	err := smtp.SendMail(client.SMTPServerURL, client.Auth, client.Username, []string{booking.Email}, message.Bytes())

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Submit",
			What:  "failure to send a booking summary mail",
			Why: errors.Meta{
				"Hotel Id":        booking.HotelID,
				"Booking Id":      booking.ID,
				"SMTP Server URL": client.SMTPServerURL,
			},
			Who: err,
		})
	}

	return nil
}
