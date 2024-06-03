package hotel

import (
	"net/http"
	"time"

	"github.com/bastean/bookingo/pkg/cmd/server/service/auth"
	"github.com/bastean/bookingo/pkg/cmd/server/service/hotel"
	"github.com/bastean/bookingo/pkg/cmd/server/util/errs"
	"github.com/bastean/bookingo/pkg/cmd/server/util/key"
	"github.com/bastean/bookingo/pkg/cmd/server/util/reply"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := new(hotel.LoginQuery)

		err := c.BindJSON(query)

		if err != nil {
			c.Error(errs.BindingJSON(err, "Login"))
			c.Abort()
			return
		}

		hotel, err := hotel.Login.Handle(query)

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		token, err := auth.GenerateJWT(auth.Payload{
			key.Exp:     time.Now().Add((24 * time.Hour) * 7).Unix(),
			key.HotelId: hotel.Id,
		})

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		session := sessions.Default(c)

		session.Set(key.Authorization, "Bearer "+token)

		session.Save()

		c.JSON(http.StatusOK, reply.JSON(true, "logged in", reply.Payload{}))
	}
}
