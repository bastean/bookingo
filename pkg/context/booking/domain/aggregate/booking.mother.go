package aggregate

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
)

func RandomBooking() *Booking {
	hotelID, _ := valueobj.RandomId()
	id, _ := valueobj.RandomId()
	firstName, _ := valueobj.RandomName()
	lastName, _ := valueobj.RandomName()
	email, _ := valueobj.RandomEmail()
	phone, _ := valueobj.RandomPhone()
	checkIn, _ := valueobj.RandomCheck()
	checkOut, _ := valueobj.RandomCheck()
	room, _ := valueobj.RandomRoom()
	currency, _ := valueobj.RandomCurrency()
	total, _ := valueobj.RandomTotal()

	booking, _ := NewBooking(
		hotelID.Value(),
		id.Value(),
		firstName.Value(),
		lastName.Value(),
		email.Value(),
		phone.Value(),
		checkIn.Value(),
		checkOut.Value(),
		room.Value(),
		currency.Value(),
		total.Value(),
	)

	return booking
}
