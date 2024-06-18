package login

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
)

func RandomQuery() *Query {
	email := valueobj.RandomEmail()
	password := valueobj.RandomPassword()

	return &Query{
		Email:    email.Value(),
		Password: password.Value(),
	}
}
