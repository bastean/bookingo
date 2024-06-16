package aggregate

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
)

func RandomHotel() *Hotel {
	id, _ := valueobj.RandomId()
	name, _ := valueobj.RandomName()
	email, _ := valueobj.RandomEmail()
	phone, _ := valueobj.RandomPhone()
	password, _ := valueobj.RandomPassword()

	hotel, _ := NewHotel(&HotelPrimitive{
		Id:       id.Value(),
		Name:     name.Value(),
		Email:    email.Value(),
		Phone:    phone.Value(),
		Password: password.Value(),
	})

	return hotel
}
