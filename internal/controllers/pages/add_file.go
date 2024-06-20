package pages

import (
	"github.com/Danila331/mifiotsos/internal/pkg"
	"github.com/labstack/echo/v4"
)

// Функция для отобрадения формы добавления файла
func AddFilePage(c echo.Context) error {
	err := pkg.HtmlPageRender("forms/addfile.html", "addfile", c)
	if err != nil {
		return err
	}
	return nil
}
