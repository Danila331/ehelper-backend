package pages

import (
	"path/filepath"
	"text/template"

	"github.com/Danila331/mifiotsos/internal/pkg"
	"github.com/labstack/echo/v4"
)

func GraphiksPageConf(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "graphicsconf.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "graphics", nil)
	pkg.GetLineGraphic()
	return nil
}
