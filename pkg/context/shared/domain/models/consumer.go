package models

import (
	"github.com/bastean/bookingo/pkg/context/shared/domain/messages"
	"github.com/bastean/bookingo/pkg/context/shared/domain/queues"
)

type Consumer interface {
	SubscribedTo() []*queues.Queue
	On(message *messages.Message) error
}
