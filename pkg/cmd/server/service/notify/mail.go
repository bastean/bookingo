package notify

import (
	"github.com/bastean/bookingo/pkg/context/notify/domain/model"
	"github.com/bastean/bookingo/pkg/context/notify/infrastructure/communication/mail"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/transports"
)

func NewMailAccountConfirmation(smtp *transports.SMTP) model.Transport {
	return &mail.AccountConfirmation{
		SMTP: smtp,
	}
}
