package pages

import (
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

func TeamPage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "team.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "team", nil)
	return nil
}
