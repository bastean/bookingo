package user

import (
	"net/http"

	"github.com/bastean/bookingo/pkg/cmd/server/service/user"
	"github.com/bastean/bookingo/pkg/cmd/server/util/errs"
	"github.com/bastean/bookingo/pkg/cmd/server/util/reply"
	"github.com/gin-gonic/gin"
)

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		command := new(user.CreateCommand)

		err := c.BindJSON(command)

		if err != nil {
			c.Error(errs.BindingJSON(err, "Create"))
			c.Abort()
			return
		}

		err = user.Create.Handle(command)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.JSON(http.StatusCreated, reply.JSON(true, "account created", reply.Payload{}))
	}
}
