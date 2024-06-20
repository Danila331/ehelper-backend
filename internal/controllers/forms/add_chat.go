package forms

import (
	"fmt"

	"github.com/Danila331/mifiotsos/internal/models"
	"github.com/Danila331/mifiotsos/internal/pkg"
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

	err = pkg.HtmlPageRender("submit/addchat_submit.html", "addchat_submit", c)
	if err != nil {
		return err
	}
	return nil
}
