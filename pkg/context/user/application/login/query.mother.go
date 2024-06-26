package login

import (
	"github.com/bastean/bookingo/pkg/context/user/domain/valueobj"
)

func RandomQuery() *Query {
	email, _ := valueobj.RandomEmail()
	password, _ := valueobj.RandomPassword()

	return &Query{
		Email:    email.Value(),
		Password: password.Value(),
	}
}
