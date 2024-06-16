package delete

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
)

func RandomCommand() *Command {
	id, _ := valueobj.RandomId()
	password, _ := valueobj.RandomPassword()

	return &Command{
		Id:       id.Value(),
		Password: password.Value(),
	}
}
