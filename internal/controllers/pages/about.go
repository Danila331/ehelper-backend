package pages

import (
	"github.com/Danila331/mifiotsos/internal/pkg"
	"github.com/labstack/echo/v4"
)

func AboutPage(c echo.Context) error {
	err := pkg.HtmlPageRender("processing.html", "processing", c)
	if err != nil {
		return err
	}
	return nil
}
