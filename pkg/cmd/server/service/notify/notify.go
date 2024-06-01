package notify

import (
	"github.com/bastean/bookingo/pkg/context/notify/application/send"
	"github.com/bastean/bookingo/pkg/context/notify/domain/model"
)

var SendAccountConfirmation = new(send.Send)

func Init(transport model.Transport) error {
	SendAccountConfirmation.Transport = transport

	return nil
}
