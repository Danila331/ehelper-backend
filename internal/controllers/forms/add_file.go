package forms

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Danila331/mifiotsos/internal/models"
	"github.com/Danila331/mifiotsos/internal/pkg"
	"github.com/labstack/echo/v4"
)

// Функция для обработки формы добваления файла
func AddFileForm(c echo.Context) error {
	var conference models.Conferences
	file, err := c.FormFile("fileToUpload")
	if err != nil {
		return err
	}

	// Открываем файл для чтения
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Создаем путь для сохранения файла на локальной машине
	uploadsDir := "uploads"
	if err := os.MkdirAll(uploadsDir, os.ModePerm); err != nil {
		return err
	}
	dstPath := filepath.Join(uploadsDir, file.Filename)

	// Создаем файл на локальной машине
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Копируем содержимое файла из запроса в файл на локальной машине
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	conference.Filepath = dstPath
	conference.Status = "no"

	fmt.Println(dstPath)
	err = conference.Create()
	fmt.Println(err)
	if err != nil {
		return err
	}

	println("File uploaded successfully")
	// Закончился код загрузки файла
	err = pkg.S3LoadFile(file.Filename, dstPath)
	if err != nil {
		return err
	}
	err = pkg.HtmlPageRender("submit/addfile_submit.html", "addfile_submit", c)
	if err != nil {
		return err
	}
	return nil
}
