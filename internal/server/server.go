package server

import (
	"github.com/Danila331/mifiotsos/internal/controllers/forms"
	"github.com/Danila331/mifiotsos/internal/controllers/midleware"
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
	// e.Use(midleware.AuthMiddleware)
	// маршруты для обычных смертных
	e.GET("/", pages.MainPage)

	sign := e.Group("/sign")
	sign.GET("/", pages.SignPage)
	sign.POST("/submit", forms.SignForm)

	login := e.Group("/login")
	login.GET("/", pages.LoginPage)
	login.POST("/submit", forms.LoginForm)

	addchat := e.Group("/add-chat")
	addchat.Use(midleware.AuthMiddleware)
	addchat.GET("/", pages.AddChatPage)
	addchat.POST("/submit", forms.AddChatForm)

	addFile := e.Group("/add-file")
	addFile.Use(midleware.AuthMiddleware)
	addFile.GET("/", pages.AddFilePage)
	addFile.POST("/submit", forms.AddFileForm)

	chat := e.Group("/chat")
	chat.Use(midleware.AuthMiddleware)
	chat.GET("/statistic", pages.StatisticPageChat)
	chat.GET("/fulstatistic", pages.FulStatisticPageChat)

	conf := e.Group("/conf")
	conf.GET("/fulstatistic", pages.FulStatisticPageConf)
	conf.GET("/statistic", pages.StatisticPageConf)
	conf.GET("/graphics", pages.GraphiksPageConf)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
