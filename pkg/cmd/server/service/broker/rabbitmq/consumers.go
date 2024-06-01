package rabbitmq

import (
	"github.com/bastean/bookingo/pkg/context/notify/application/send"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/queues"
)

var NotifySendAccountConfirmationQueueConsumer = &send.CreatedSucceededEventConsumer{
	Queues: []*queues.Queue{NotifySendAccountConfirmationQueue},
}

func Consumers(send *send.Send) []models.Consumer {
	NotifySendAccountConfirmationQueueConsumer.UseCase = send

	return []models.Consumer{
		NotifySendAccountConfirmationQueueConsumer,
	}
}
