package pages

import (
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

// Фукнция для отображения главной страницы
func MainPage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "main.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "main", nil)
	return nil
}
