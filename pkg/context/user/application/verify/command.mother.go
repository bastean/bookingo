package verify

import (
	"github.com/bastean/bookingo/pkg/context/user/domain/valueobj"
)

func RandomCommand() *Command {
	id, _ := valueobj.RandomId()

	return &Command{
		Id: id.Value(),
	}
}
