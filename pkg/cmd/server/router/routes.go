package router

import (
	"github.com/bastean/bookingo/pkg/cmd/server/handler/hotel"
	"github.com/bastean/bookingo/pkg/cmd/server/handler/page"
	"github.com/bastean/bookingo/pkg/cmd/server/middleware"
)

func InitRoutes() {
	router.NoRoute(page.Default())

	public := router.Group("/")

	public.GET("/", page.Home())
	public.PUT("/", hotel.Create())
	public.POST("/", hotel.Login())

	public.GET("/verify/:id", hotel.Verify())

	auth := public.Group("/dashboard", middleware.VerifyAuthentication())

	auth.GET("/", page.Dashboard())
	auth.PATCH("/", hotel.Update())
	auth.DELETE("/", hotel.Delete())
}
