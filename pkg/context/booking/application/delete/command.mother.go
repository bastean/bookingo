package delete

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
)

func RandomCommand() *Command {
	hotelID, _ := valueobj.RandomId()
	id, _ := valueobj.RandomId()

	return &Command{
		HotelID: hotelID.Value(),
		ID:      id.Value(),
	}
}
