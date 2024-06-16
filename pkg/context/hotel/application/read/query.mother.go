package read

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
)

func RandomQuery() *Query {
	id, _ := valueobj.RandomId()

	return &Query{
		Id: id.Value(),
	}
}
