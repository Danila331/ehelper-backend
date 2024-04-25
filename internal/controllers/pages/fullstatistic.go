package pages

import (
	"path/filepath"
	"text/template"

	"github.com/Danila331/mifiotsos/internal/pkg"
	"github.com/labstack/echo/v4"
)

func FulStatisticPageConf(c echo.Context) error {

	fullstatistic, err := pkg.GetFulResultConf()
	if err != nil {
		return err
	}

	htmlFiles := []string{
		filepath.Join("./", "templates", "fulstatisticconf.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "fulstatistic", fullstatistic)
	return nil
}

// func FulStatisticPageChat(c echo.Context) error {

// 	fullstatistic, err := pkg.GetFulResultCont()
// 	if err != nil {
// 		return err
// 	}

// 	htmlFiles := []string{
// 		filepath.Join("./", "templates", "fulstatisticconf.html"),
// 	}

// 	templ, err := template.ParseFiles(htmlFiles...)
// 	if err != nil {
// 		return err
// 	}

// 	templ.ExecuteTemplate(c.Response(), "fulstatistic", fullstatistic)
// 	return nil
// }
