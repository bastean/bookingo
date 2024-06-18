package delete

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
)

func RandomCommand() *Command {
	id := valueobj.RandomId()
	password := valueobj.RandomPassword()

	return &Command{
		Id:       id.Value(),
		Password: password.Value(),
	}
}
