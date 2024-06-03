package smtp

import (
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/transports"
)

func New(host, port, hotelname, password, serverURL string) *transports.SMTP {
	return transports.NewSMTP(host, port, hotelname, password, serverURL)
}
