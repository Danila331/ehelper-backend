package server

import (
	"github.com/Danila331/mifiotsos/internal/controllers/forms"
	"github.com/Danila331/mifiotsos/internal/controllers/pages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "./static")

	e.GET("/", pages.AddFilePage)
	e.POST("/add-file_submit", forms.AddFileForm)

	e.GET("/sign", pages.SignPage)
	e.POST("/sign_submit", forms.SignForm)

	e.GET("/login", pages.LoginPage)
	e.POST("/login_submit", forms.LoginForm)

	e.GET("/chat/statistic", pages.StatisticPageChat)
	// e.GET("/chat/fulstatistic", pages.FulStatisticPageChat)
	e.GET("/conf/fulstatistic", pages.FulStatisticPageConf)
	e.GET("/conf/statistic", pages.StatisticPageConf)
	e.GET("/conf/graphics", pages.GraphiksPageConf)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
