FROM golang:latest

#создали директорию 
RUN mkdir /awesomeproject4
#перенесли в неё проект
COPY  /awesomeProject4/ /awesomeproject4/
#сделали директорию рабочей
WORKDIR /awesomeproject4

#скачиваем зависимости
RUN go get -u "github.com/gin-gonic/gin"
RUN go get -u "github.com/go-sql-driver/mysql"
RUN go get -u "github.com/jmoiron/sqlx"

# Соберём приложение 
RUN go build -o main .

# Запустим приложение
CMD ["/awesomeproject4/main"]
#"/awesomeproject4/main" - путь до бинаря