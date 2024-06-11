package terminal_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/bastean/bookingo/pkg/context/hotel/domain/event"
	"github.com/bastean/bookingo/pkg/context/hotel/infrastructure/communication/terminal"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/loggers"
	"github.com/stretchr/testify/suite"
)

type HotelTerminalAccountConfirmationTransportTestSuite struct {
	suite.Suite
	sut       models.Transport
	logger    *loggers.LoggerMock
	serverURL string
}

func (suite *HotelTerminalAccountConfirmationTransportTestSuite) SetupTest() {
	suite.logger = new(loggers.LoggerMock)

	suite.serverURL = os.Getenv("URL")

	suite.sut = &terminal.AccountConfirmation{
		Logger:    suite.logger,
		ServerURL: suite.serverURL,
	}
}

func (suite *HotelTerminalAccountConfirmationTransportTestSuite) TestSubmit() {
	message := event.RandomCreatedSucceeded()

	attributes := new(event.CreatedSucceededAttributes)

	json.Unmarshal(message.Attributes, attributes)

	confirmationLink := fmt.Sprintf("Hi %s, please confirm your account through this link: %s/verify/%s", attributes.Name, suite.serverURL, attributes.ID)

	suite.logger.Mock.On("Info", confirmationLink)

	suite.NoError(suite.sut.Submit(attributes))

	suite.logger.AssertExpectations(suite.T())
}

func TestIntegrationHotelTerminalAccountConfirmationTransportSuite(t *testing.T) {
	suite.Run(t, new(HotelTerminalAccountConfirmationTransportTestSuite))
}
