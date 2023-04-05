package pages

import (
	"awesomeProject4/datamysql"
	"awesomeProject4/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

func Project_page(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		// Get data from DB
		projectsData := datamysql.ExtractData_Projects()
		returningResult := gin.H{
			"projects": projectsData,
		}
		admin, _ := c.Cookie("admin")

		if admin == "true" {
			returningResult["admin"] = true
		}
		c.HTML(200, "projects.html", returningResult)
	} else {

		c.Redirect(http.StatusFound, "/login")
	}
}

func Project_info(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		id := c.Param("id")
		project_info := datamysql.ExtractDataProject(id)

		if reflect.DeepEqual(project_info, model.Task{}) {
			c.Redirect(http.StatusFound, "/")
			return
		}

		tasks := datamysql.ExtractDataProject_info(id) //добавил для отрисовки
		date := time.Now()
		for i := 0; i < len(tasks); i++ {
			tasks[i].Url = "../project/" + string(id) + "/task/" + string(tasks[i].Id)
			finTime, _ := time.Parse("2006-01-02", tasks[i].End)
			tasks[i].Ongoing = finTime.Before(date)
			fmt.Println(tasks[i].Ongoing)
		}

		fmt.Println(time.Now())

		returningResult := gin.H{
			"tasks":         tasks,
			"project_title": project_info.Title,
			"project_start": project_info.Start,
			"project_end":   project_info.End,
			"project_desc":  project_info.Text,
			"id":            id,
		}

		admin, _ := c.Cookie("admin")

		if admin == "true" {
			returningResult["admin"] = true
		}

		c.HTML(200, "project_info.html", returningResult)
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

func All_calendar(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		// Get data from DB
		returningResult := gin.H{}
		admin, _ := c.Cookie("admin")

		if admin == "true" {
			returningResult["admin"] = true
		}
		c.HTML(200, "all_calendar.html", returningResult)
	} else {
		c.Redirect(http.StatusFound, "/login")
	}

}
func Project_calendar(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		// Get data from DB

		id := c.Param("id")

		project_info := datamysql.ExtractDataProject(id)

		if reflect.DeepEqual(project_info, model.Task{}) {
			c.Redirect(http.StatusFound, "/")
			return
		}

		returningResult := gin.H{
			"Id": id,
		}
		admin, _ := c.Cookie("admin")

		if admin == "true" {
			returningResult["admin"] = true
		}

		c.HTML(200, "calendar.html", returningResult)
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

func Login_page(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		c.Redirect(http.StatusFound, "/")
	}
	c.HTML(200, "login.html", nil)
}

func Registration(c *gin.Context) {
	var user *model.Registr
	decode := json.NewDecoder(c.Request.Body).Decode(&user)
	is_logining := datamysql.ExtractData(datamysql.Db, string(user.Login), string(user.Pass))
	user_id := strconv.Itoa(is_logining.Id)

	if is_logining.Login == true { //отправка подтверждения логина
		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie("user", user_id, 3600, "/", "localhost", true, false)
		returningResult := gin.H{}
		if is_logining.Admin {
			returningResult["admin"] = true
			c.SetCookie("admin", string("true"), 3600, "/", "localhost", true, false)
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error": "Неверный логин или пароль.",
		})
		return
	}

	c.JSON(http.StatusOK, nil)

	if decode != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
		return
	}

}
func Task_info(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		taskId := c.Param("task_id")
		id := c.Param("id")
		all_persons := datamysql.ExtractDataUsers()
		task := datamysql.ExtractDataTask(taskId)
		persons := datamysql.ExtractDataUsers_Task(taskId)
		if reflect.DeepEqual(task, model.Task{}) {
			c.Redirect(http.StatusFound, "/project/"+id)
			return
		}
		returningResult := gin.H{
			"title":       task.Title,
			"startDate":   task.Start,
			"endDate":     task.End,
			"text":        task.Text,
			"person":      persons,
			"all_persons": all_persons,
			"id":          taskId,
			"Is_finished": task.Is_finished,
		}
		admin, _ := c.Cookie("admin")

		if admin == "true" {
			returningResult["admin"] = true
		}

		c.HTML(200, "task_info.html", returningResult)
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

func Edit_info(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		id := c.Param("id")
		admin, _ := c.Cookie("admin")
		project_info := datamysql.ExtractDataProject(id)
		fmt.Println(project_info)
		if reflect.DeepEqual(project_info, model.Task{}) {
			c.Redirect(http.StatusFound, "/")
			return
		}

		returningResult := gin.H{
			"title":      project_info.Title,
			"text":       project_info.Text,
			"id":         id,
			"color":      project_info.BackgroundColor,
			"text_color": project_info.TextColor,
		}
		if admin == "true" {
			returningResult["admin"] = true
		}
		c.HTML(200, "edit_info.html", returningResult)
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}
func Create_project(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		returningResult := gin.H{}
		admin, _ := c.Cookie("admin")

		if admin == "true" {
			returningResult["admin"] = true
		}
		c.HTML(200, "create_project.html", returningResult)
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}
func Create_task(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		all_persons := datamysql.ExtractDataUsers()

		id := c.Param("id")
		returningResult := gin.H{
			"all_persons": all_persons,
			"id":          id,
		}
		admin, _ := c.Cookie("admin")
		if admin == "true" {
			returningResult["admin"] = true
		}
		c.HTML(200, "create_task.html", returningResult)
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}
func Get_all_calendar(c *gin.Context) {
	tasks := datamysql.ExtractDataProject_info_ALL()
	for i := 0; i < len(tasks); i++ {
		tasks[i].Url = "../project/" + string(tasks[i].Project_id) + "/task/" + string(tasks[i].Id)
	}

	c.JSON(http.StatusOK, tasks)
}

func Get_tasks(c *gin.Context) {
	id := c.Param("id")
	tasks := datamysql.ExtractDataProject_info(id)
	for i := 0; i < len(tasks); i++ {
		tasks[i].Url = "../project/" + string(id) + "/task/" + string(tasks[i].Id)
	}
	c.JSON(http.StatusOK, tasks)
	return

}
func Person_tasks(c *gin.Context) {
	user, err := c.Cookie("user")
	if err == nil {
		tasks := datamysql.ExtractDataProject_and_Task(user)
		for i := 0; i < len(tasks); i++ {
			tasks[i].Url = "../project/" + string(tasks[i].Project_id) + "/task/" + string(tasks[i].Id)
		}
		returningResult := gin.H{"tasks": tasks}
		admin, _ := c.Cookie("admin")
		if admin == "true" {
			returningResult["admin"] = true
		}
		c.HTML(200, "all_tasks.html", returningResult)
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

func NewPerson(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		returningResult := gin.H{}
		admin, _ := c.Cookie("admin")
		if admin == "true" {
			returningResult["admin"] = true
			c.HTML(200, "create_login.html", returningResult)
		} else {
			c.HTML(200, "projects.html", nil)
		}

	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}
func CreateUser(c *gin.Context) {
	var user *model.CreateUser
	decode := json.NewDecoder(c.Request.Body).Decode(&user)
	fmt.Println(user.FName, " ", user.LName, " ", user.Login, " ", user.Password)
	if decode != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
	}
	isCreated := datamysql.AddPerson(user.Login, user.Password, user.FName, user.LName)

	if isCreated {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error": "Пользователь с данным логином существует.",
		})
	}
	return
}
func CreateProject(c *gin.Context) {
	var project *model.CreateProject
	decode := json.NewDecoder(c.Request.Body).Decode(&project)
	fmt.Println(project.Name, " ", project.Description, " ", project.Colour)
	if decode != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
	}
	datamysql.AddProject(project.Name, project.Description, project.Colour, project.TextColor)
}

func CreateTask(c *gin.Context) {
	var task *model.Task
	id := c.Param("id")
	decode := json.NewDecoder(c.Request.Body).Decode(&task)
	//fmt.Println(project.Name, " ", project.Description, " ", project.Colour)
	if decode != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
	}
	datamysql.AddTask(task.Title, task.Start, task.End, task.Text, id, task.Person_Mas)

}

func Update_project(c *gin.Context) {
	// update data base AHAHAHHAHAHAA
	id := c.Param("id")
	var project *model.CreateProject
	decode := json.NewDecoder(c.Request.Body).Decode(&project)
	fmt.Println(project.Name, " ", project.Description, " ", project.Colour)
	if decode != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
	}
	datamysql.UpdateProject(id, project.Name, project.Description, project.Colour, project.TextColor)
}

func Update_task(c *gin.Context) {
	var task *model.Task
	id := c.Param("task_id")
	decode := json.NewDecoder(c.Request.Body).Decode(&task)
	fmt.Println(id)
	if decode != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
	}
	datamysql.UpdateTask(id, task.Title, task.Start, task.End, task.Text, task.Person_Mas, task.Is_finished)

}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	task_id := c.Param("task_id")

	datamysql.DeleteTaskFromDB(id, task_id)
}

func DeleteProject(c *gin.Context) {
	id := c.Param("id")

	datamysql.DeleteProjectFromDB(id)
}
func AllPersons(c *gin.Context) {
	_, err := c.Cookie("user")
	if err == nil {
		// Get data from DB
		usersData := datamysql.GetNotAdminUsers()
		returningResult := gin.H{
			"users": usersData,
		}
		admin, _ := c.Cookie("admin")

		if admin == "true" {
			returningResult["admin"] = true
		}
		c.HTML(200, "all_users.html", returningResult)
	} else {

		c.Redirect(http.StatusFound, "/login")
	}
}

func DeleteUser(c *gin.Context) {
	user_id := c.Param("id")

	datamysql.DelUser(user_id)
}
