package aggregate

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
)

func RandomHotel() *Hotel {
	id := valueobj.RandomId()
	name := valueobj.RandomName()
	email := valueobj.RandomEmail()
	phone := valueobj.RandomPhone()
	password := valueobj.RandomPassword()

	hotel, err := NewHotel(&HotelPrimitive{
		Id:       id.Value(),
		Name:     name.Value(),
		Email:    email.Value(),
		Phone:    phone.Value(),
		Password: password.Value(),
	})

	if err != nil {
		errors.Panic(err.Error(), "RandomHotel")
	}

	return hotel
}
