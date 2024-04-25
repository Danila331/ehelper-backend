package pages

import (
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

func AddFilePage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "addfile.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "addfile", nil)
	return nil
}
