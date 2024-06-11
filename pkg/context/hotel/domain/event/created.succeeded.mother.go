package event

import (
	"github.com/bastean/bookingo/pkg/context/hotel/domain/valueobj"
	"github.com/bastean/bookingo/pkg/context/shared/domain/messages"
)

func RandomCreatedSucceeded() *messages.Message {
	id, _ := valueobj.RandomId()
	name, _ := valueobj.RandomName()
	email, _ := valueobj.RandomEmail()
	phone, _ := valueobj.RandomPhone()

	event, _ := NewCreatedSucceeded(&CreatedSucceeded{
		Attributes: &CreatedSucceededAttributes{
			ID:    id.Value(),
			Name:  name.Value(),
			Email: email.Value(),
			Phone: phone.Value(),
		},
	})

	return event
}
