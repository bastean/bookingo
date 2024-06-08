package read

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
)

func RandomQuery() *Query {
	hotelID, _ := valueobj.RandomId()
	ID, _ := valueobj.RandomId()

	return &Query{
		HotelID: hotelID.Value(),
		ID:      ID.Value(),
	}
}
