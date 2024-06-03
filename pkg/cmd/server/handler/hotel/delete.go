package hotel

import (
	"net/http"

	"github.com/bastean/bookingo/pkg/cmd/server/service/hotel"
	"github.com/bastean/bookingo/pkg/cmd/server/util/errs"
	"github.com/bastean/bookingo/pkg/cmd/server/util/key"
	"github.com/bastean/bookingo/pkg/cmd/server/util/reply"
	"github.com/gin-gonic/gin"
)

func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get(key.HotelId)

		if !exists {
			c.Error(errs.MissingKey(key.HotelId, "Delete"))
			c.Abort()
			return
		}

		command := new(hotel.DeleteCommand)

		err := c.BindJSON(command)

		if err != nil {
			c.Error(errs.BindingJSON(err, "Delete"))
			c.Abort()
			return
		}

		command.Id = id.(string)

		err = hotel.Delete.Handle(command)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, reply.JSON(true, "account deleted", reply.Payload{}))
	}
}
