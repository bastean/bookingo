package terminal_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bastean/bookingo/pkg/context/booking/domain/event"
	"github.com/bastean/bookingo/pkg/context/booking/infrastructure/communication/terminal"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/loggers"
	"github.com/stretchr/testify/suite"
)

type BookingTerminalSummaryTransportTestSuite struct {
	suite.Suite
	sut    models.Transport
	logger *loggers.LoggerMock
}

func (suite *BookingTerminalSummaryTransportTestSuite) SetupTest() {
	suite.logger = new(loggers.LoggerMock)

	suite.sut = &terminal.Summary{
		Logger: suite.logger,
	}
}

func (suite *BookingTerminalSummaryTransportTestSuite) TestSubmit() {
	message := event.RandomCreatedSucceeded()

	booking := new(event.CreatedSucceededAttributes)

	json.Unmarshal(message.Attributes, booking)

	summary := fmt.Sprintf(
		`
		Hi %s %s, 
		
		Here is the summary of your booking:

		Check-in: %s
		Check-out: %s
		Room: %s

		Total(%s): %s
		`,
		booking.FirstName, booking.LastName,
		booking.CheckIn, booking.CheckOut,
		booking.Room,
		booking.Currency, booking.Total,
	)

	suite.logger.Mock.On("Info", summary)

	suite.NoError(suite.sut.Submit(booking))

	suite.logger.AssertExpectations(suite.T())
}

func TestIntegrationBookingTerminalSummaryTransportSuite(t *testing.T) {
	suite.Run(t, new(BookingTerminalSummaryTransportTestSuite))
}
