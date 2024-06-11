package send_test

import (
	"encoding/json"
	"testing"

	"github.com/bastean/bookingo/pkg/context/hotel/application/send"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/event"
	"github.com/bastean/bookingo/pkg/context/hotel/infrastructure/communication"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/domain/queues"
	"github.com/bastean/bookingo/pkg/context/shared/domain/types"
	"github.com/stretchr/testify/suite"
)

type CreatedSucceededEventConsumerTestSuite struct {
	suite.Suite
	sut       models.Consumer
	useCase   models.UseCase[*event.CreatedSucceeded, types.Empty]
	transport *communication.TransportMock
	queues    []*queues.Queue
}

func (suite *CreatedSucceededEventConsumerTestSuite) SetupTest() {
	queueName := queues.NewQueueName(&queues.QueueName{
		Module: "queue",
		Action: "assert",
		Event:  "test.succeeded",
	})

	suite.queues = append(suite.queues, &queues.Queue{
		Name: queueName,
	})

	suite.transport = new(communication.TransportMock)

	suite.useCase = &send.Send{
		Transport: suite.transport,
	}

	suite.sut = &send.CreatedSucceededEventConsumer{
		UseCase: suite.useCase,
		Queues:  suite.queues,
	}
}

func (suite *CreatedSucceededEventConsumerTestSuite) TestEventConsumer() {
	message := event.RandomCreatedSucceeded()

	attributes := new(event.CreatedSucceededAttributes)

	json.Unmarshal(message.Attributes, attributes)

	suite.transport.On("Submit", attributes)

	suite.NoError(suite.sut.On(message))

	suite.transport.AssertExpectations(suite.T())
}

func TestUnitCreatedSucceededEventConsumerSuite(t *testing.T) {
	suite.Run(t, new(CreatedSucceededEventConsumerTestSuite))
}
