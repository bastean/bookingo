package delete

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
)

func RandomCommand() *Command {
	hotelId := valueobj.RandomId()
	id := valueobj.RandomId()

	return &Command{
		HotelId: hotelId.Value(),
		Id:      id.Value(),
	}
}
