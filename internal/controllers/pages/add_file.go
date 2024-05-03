package pages

import (
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

// Функция для отобрадения формы добавления файла
func AddFilePage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "forms", "addfile.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "addfile", nil)
	return nil
}
