package update

import (
	"github.com/bastean/bookingo/pkg/context/user/domain/valueobj"
)

func RandomCommand() *Command {
	id, _ := valueobj.RandomId()
	email, _ := valueobj.RandomEmail()
	username, _ := valueobj.RandomUsername()
	password, _ := valueobj.RandomPassword()
	updatedPassword, _ := valueobj.RandomPassword()

	return &Command{
		Id:              id.Value(),
		Email:           email.Value(),
		Username:        username.Value(),
		Password:        password.Value(),
		UpdatedPassword: updatedPassword.Value(),
	}
}
