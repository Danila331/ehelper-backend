package pages

import (
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

func AddChatPage(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("./", "templates", "forms", "addchat.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "addchat", nil)
	return nil
}
