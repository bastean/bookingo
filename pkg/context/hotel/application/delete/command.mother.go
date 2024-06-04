package delete

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
)

func RandomCommand() *Command {
	id, _ := valueobj.RandomId()
	password, _ := valueobj.RandomPassword()

	return &Command{
		ID:       id.Value(),
		Password: password.Value(),
	}
}
