package hotel

import (
	"github.com/bastean/bookingo/pkg/context/hotel/application/create"
	"github.com/bastean/bookingo/pkg/context/hotel/application/delete"
	"github.com/bastean/bookingo/pkg/context/hotel/application/login"
	"github.com/bastean/bookingo/pkg/context/hotel/application/read"
	"github.com/bastean/bookingo/pkg/context/hotel/application/update"
	"github.com/bastean/bookingo/pkg/context/hotel/application/verify"
	"github.com/bastean/bookingo/pkg/context/hotel/domain/model"
	"github.com/bastean/bookingo/pkg/context/shared/domain/models"
)

var Create *create.CommandHandler

var Read *read.QueryHandler

var Update *update.CommandHandler

var Delete *delete.CommandHandler

var Verify *verify.CommandHandler

var Login *login.QueryHandler

func Init(repository model.Repository, broker models.Broker, hashing models.Hashing) {
	Create = NewCreate(repository, broker)

	Read = NewRead(repository)

	Update = NewUpdate(repository, hashing)

	Delete = NewDelete(repository, hashing)

	Verify = NewVerify(repository)

	Login = NewLogin(repository, hashing)
}
