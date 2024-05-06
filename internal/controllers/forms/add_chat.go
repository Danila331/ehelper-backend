package forms

import (
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/Danila331/mifiotsos/internal/models"
	"github.com/labstack/echo/v4"
)

// Функция для обработки формы добавления чата
func AddChatForm(c echo.Context) error {
	var user models.User

	email := c.Get("email").(string)
	user, err := user.ReadByEmail(email)

	if err != nil {
		return err
	}

	chatid := c.FormValue("chat")

	user.ChatsId = fmt.Sprintf("%s %s", user.ChatsId, chatid)

	err = user.Update()
	if err != nil {
		return err
	}

	htmlFiles := []string{
		filepath.Join("./", "templates", "submit", "addchat_submit.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "addchat_submit", nil)

	return nil
}
