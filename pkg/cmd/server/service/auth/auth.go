package auth

import (
	"github.com/bastean/bookingo/pkg/cmd/server/service/env"
	"github.com/bastean/bookingo/pkg/context/shared/infrastructure/authentications"
)

type Payload authentications.Payload

var auth = authentications.NewAuthentication(env.JWT.SecretKey)

var GenerateJWT = auth.GenerateJWT

var ValidateJWT = auth.ValidateJWT
