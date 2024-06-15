# Используем официальный образ Golang в качестве базового образа
FROM golang:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /echo-todo

# Копируем файлы go.mod и go.sum в рабочую директорию
COPY go.mod go.sum ./

# Загружаем зависимости Go
RUN go mod download

# Копируем все файлы исходного кода в рабочую директорию
COPY . .

COPY config/config.yml ./config/config.yml

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# Собираем Go-приложение
RUN go build -o echo-todo ./cmd

RUN ls -l ./config

# Указываем, что контейнер слушает на порту 8080
EXPOSE 8080

# Команда, которая будет выполняться при запуске контейнера
CMD ["./echo-todo"]
