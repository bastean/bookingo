package event

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/messages"
)

func RandomCreatedSucceeded() *messages.Message {
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

	event, _ := NewCreatedSucceeded(&CreatedSucceeded{
		Attributes: &CreatedSucceededAttributes{
			HotelID:   hotelID.Value(),
			ID:        id.Value(),
			FirstName: firstName.Value(),
			LastName:  lastName.Value(),
			Email:     email.Value(),
			Phone:     phone.Value(),
			CheckIn:   checkIn.Value(),
			CheckOut:  checkOut.Value(),
			Room:      room.Value(),
			Currency:  currency.Value(),
			Total:     total.Value(),
		},
	})

	return event
}
