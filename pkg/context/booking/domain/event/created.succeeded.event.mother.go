package event

import (
	"github.com/bastean/bookingo/pkg/context/booking/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/messages"
)

func RandomCreatedSucceeded() *messages.Message {
	hotelId := valueobj.RandomId()
	id := valueobj.RandomId()
	firstName := valueobj.RandomName()
	lastName := valueobj.RandomName()
	email := valueobj.RandomEmail()
	phone := valueobj.RandomPhone()
	checkIn := valueobj.RandomCheck()
	checkOut := valueobj.RandomCheck()
	room := valueobj.RandomRoom()
	currency := valueobj.RandomCurrency()
	total := valueobj.RandomTotal()

	event, err := NewCreatedSucceeded(&CreatedSucceeded{
		Attributes: &CreatedSucceededAttributes{
			HotelId:   hotelId.Value(),
			Id:        id.Value(),
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

	if err != nil {
		errors.Panic(err.Error(), "RandomCreatedSucceeded")
	}

	return event
}
