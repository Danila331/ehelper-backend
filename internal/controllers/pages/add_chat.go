package pages

import (
	"github.com/Danila331/mifiotsos/internal/pkg"
	"github.com/labstack/echo/v4"
)

// Функция для отображения формы добавления чата
func AddChatPage(c echo.Context) error {
	err := pkg.HtmlPageRender("forms/addchat.html", "addchat", c)
	if err != nil {
		return err
	}
	return nil
}
