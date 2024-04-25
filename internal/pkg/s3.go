package pkg

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func S3LoadFile(filePath string) error {
	// Указываем регион, ключ доступа и секретный ключ
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"), // Замените на нужный регион
		Credentials: credentials.NewStaticCredentials("YOUR_ACCESS_KEY_ID", "YOUR_SECRET_ACCESS_KEY", ""),
	})

	if err != nil {
		return err
	}

	// Создаем новый экземпляр сервиса S3
	svc := s3.New(sess)

	// Указываем имя бакета и путь к файлу, который хотим загрузить
	bucketName := "YOUR_BUCKET_NAME" // Путь к вашему локальному файлу

	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Определяем имя файла в S3 (можно использовать имя файла из локальной системы)
	key := file.Name()

	// Выполняем загрузку файла в S3
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return err
	}

	log.Printf("File uploaded successfully to bucket: %s", bucketName)
	return nil
}
