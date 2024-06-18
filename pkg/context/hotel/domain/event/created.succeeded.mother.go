package event

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/errors"
	"github.com/bastean/bookingo/pkg/context/shared/domain/messages"
)

func RandomCreatedSucceeded() *messages.Message {
	id := valueobj.RandomId()
	name := valueobj.RandomName()
	email := valueobj.RandomEmail()
	phone := valueobj.RandomPhone()

	event, err := NewCreatedSucceeded(&CreatedSucceeded{
		Attributes: &CreatedSucceededAttributes{
			Id:    id.Value(),
			Name:  name.Value(),
			Email: email.Value(),
			Phone: phone.Value(),
		},
	})

	if err != nil {
		errors.Panic(err.Error(), "RandomCreatedSucceeded")
	}

	return event
}
