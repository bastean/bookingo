package delete

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
)

func RandomCommand() *Command {
	hotelId, _ := valueobj.RandomId()
	id, _ := valueobj.RandomId()

	return &Command{
		HotelId: hotelId.Value(),
		Id:      id.Value(),
	}
}
