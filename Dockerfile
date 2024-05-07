# # Используем образ Golang для сборки приложения
# FROM golang:latest AS builder

# # Установка рабочей директории внутри контейнера
# WORKDIR /app

# # Копируем файлы go.mod и go.sum для загрузки зависимостей
# COPY go.mod .
# COPY go.sum .

# # Загрузка зависимостей с помощью go mod download
# RUN go mod download

# # Копируем исходный код проекта в контейнер
# COPY . .

# # Переходим в папку internal
# WORKDIR /app/internal

# EXPOSE 80
# # Запускаем main.go
# CMD ["go", "run", "main.go"]





# Используем образ Golang для сборки приложения
FROM golang:latest AS builder

# Установка рабочей директории внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum для загрузки зависимостей
COPY go.mod .
COPY go.sum .

# Загрузка зависимостей с помощью go mod download
RUN go mod download

# Копируем исходный код проекта в контейнер
COPY . .

# Переходим в папку cmd
WORKDIR /app/internal

# Компилируем приложение
RUN go build -o main .

# Указываем порт, который будет использоваться приложением
EXPOSE 80

# Запускаем скомпилированное приложение
CMD ["./main"]