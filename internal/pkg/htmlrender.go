package pkg

import (
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
)

func HtmlPageRender(htmlFilePathString, htmlFileHeader string, c echo.Context) error {

	htmlFiles := []string{
		filepath.Join("./", "templates", htmlFilePathString),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), htmlFileHeader, nil)

	return nil
}
