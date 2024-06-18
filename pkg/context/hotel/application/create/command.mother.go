package create

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
)

func RandomCommand() *Command {
	id := valueobj.RandomId()
	name := valueobj.RandomName()
	email := valueobj.RandomEmail()
	phone := valueobj.RandomPhone()
	password := valueobj.RandomPassword()

	return &Command{
		Id:       id.Value(),
		Name:     name.Value(),
		Email:    email.Value(),
		Phone:    phone.Value(),
		Password: password.Value(),
	}
}
