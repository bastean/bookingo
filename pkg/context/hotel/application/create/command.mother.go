package create

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
)

func RandomCommand() *Command {
	id, _ := valueobj.RandomId()
	email, _ := valueobj.RandomEmail()
	hotelname, _ := valueobj.RandomHotelname()
	password, _ := valueobj.RandomPassword()

	return &Command{
		Id:        id.Value(),
		Email:     email.Value(),
		Hotelname: hotelname.Value(),
		Password:  password.Value(),
	}
}
