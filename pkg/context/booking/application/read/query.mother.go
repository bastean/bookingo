package read

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
)

func RandomQuery() *Query {
	hotelId := valueobj.RandomId()
	Id := valueobj.RandomId()

	return &Query{
		HotelId: hotelId.Value(),
		Id:      Id.Value(),
	}
}
