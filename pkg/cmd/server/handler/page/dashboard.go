package page

import (
	"net/http"

	"github.com/bastean/bookingo/pkg/cmd/server/component/page/dashboard"
	"github.com/bastean/bookingo/pkg/cmd/server/service/hotel"
	"github.com/bastean/bookingo/pkg/cmd/server/util/errs"
	"github.com/bastean/bookingo/pkg/cmd/server/util/key"
	"github.com/gin-gonic/gin"
)

func Dashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exists := c.Get(key.HotelId)

		if !exists {
			c.Error(errs.MissingKey(key.HotelId, "Dashboard"))
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}

		query := new(hotel.ReadQuery)

		query.Id = id.(string)

		hotel, err := hotel.Read.Handle(query)

		if err != nil {
			c.Error(err)
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}

		dashboard.Page(hotel).Render(c.Request.Context(), c.Writer)
	}
}
