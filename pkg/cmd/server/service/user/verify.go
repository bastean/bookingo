package user

import (
	"github.com/bastean/bookingo/pkg/context/user/application/verify"
	"github.com/bastean/bookingo/pkg/context/user/domain/model"
)

type VerifyCommand = verify.Command

func NewVerify(repository model.Repository) *verify.CommandHandler {
	useCase := &verify.Verify{
		Repository: repository,
	}

	return &verify.CommandHandler{
		UseCase: useCase,
	}
}
