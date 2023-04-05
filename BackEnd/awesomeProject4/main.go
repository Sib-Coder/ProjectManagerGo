package main

import (
	"awesomeProject4/datamysql"
	"awesomeProject4/pages"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var router *gin.Engine

func main() {
	//подключение к БД
	datamysql.Conect()

	router := gin.Default()

	router.LoadHTMLGlob("html/*.html")   //шаблоны
	router.Static("/static", "./static") //css, js ...
	//роуты надо прописать в html пути со "static/"

	router.GET("/", pages.Project_page)            // работает всё ок
	router.GET("/login", pages.Login_page)         // работает всё ок
	router.GET("/calendar", pages.All_calendar)    // Рома где Блять calendar.css , а так всё ок
	router.GET("/project/:id", pages.Project_info) // Рома где Блять info.css
	router.GET("/calendar/:id", pages.Project_calendar)
	router.GET("/project/:id/task/:task_id", pages.Task_info)
	router.GET("/edit/:id/", pages.Edit_info)
	router.GET("/edit/", pages.Create_project)
	router.GET("/project/:id/task/", pages.Create_task)
	router.GET("/person_tasks", pages.Person_tasks)
	router.GET("/new_person", pages.NewPerson)
	router.GET("/all_users", pages.AllPersons)
	//router.GET("/project/", )
	//api
	router.POST("/api/login", pages.Registration) // работает всё ок
	router.GET("/api/tasks", pages.Get_all_calendar)
	router.GET("/api/tasks/:id", pages.Get_tasks)
	router.POST("/api/project/:id/task/:task_id", pages.Update_task)
	router.POST("/api/project/:id", pages.Update_project)
	router.POST("/api/new_login", pages.CreateUser)
	router.POST("/api/new_project", pages.CreateProject)
	router.POST("/api/project/:id/add_task/", pages.CreateTask)
	router.POST("/api/delete_project/:id", pages.DeleteProject)
	router.POST("/api/from_project/:id/del_task/:task_id", pages.DeleteTask)
	router.POST("/api/delete_user/:id", pages.DeleteUser)

	router.Run(":3001")

}
