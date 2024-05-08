package pages

import (
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

func PricingPage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "processing.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "processing", nil)
	return nil
}
