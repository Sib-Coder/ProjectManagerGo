package datamysql

import (
	"awesomeProject4/model"
	"awesomeProject4/privat_info"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func Conect() {
	db, err := sqlx.Open("mysql", privat_info.DataBaseKey)
	if err != nil {
		panic(err)
	}
	Db = db
	//defer Db.Close() //Закрытие БД
	fmt.Println("DataBase_is_WORK")
}

func ExtractData(db *sqlx.DB, login string, password string) model.Is_login { //получение пользователя из бд
	var u model.User_DB
	res, err := db.Query("SELECT `first_name`,`last_name`,`id`, `is_admin` FROM `employees` WHERE `login`= ? AND `password` = ?;", login, password)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.FName, &u.LName, &u.Id, &u.Admin)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
	}
	result := model.Is_login{
		Login: false,
		Admin: false,
	}
	fmt.Println(u.Id)
	if u.LName != "" && u.FName != "" {
		//privat_info.Login = true
		result.Login = true
		result.Id = u.Id
		if u.Admin == 1 {
			result.Admin = true
		}
	} else {
		//privat_info.Login = false
	}

	return result
	//fmt.Println(u.LName, " ", u.FName) //пример как вырывать параметры из запроса

}

func DelUser(id string) {
	_, err := Db.Exec("DELETE FROM task_for_emp WHERE emp_id = ? ", id)
	if err != nil {
		panic(err)
	}

	_, err1 := Db.Exec("delete from employees where id = ?", id)
	if err1 != nil {
		panic(err)
	}

}
func ExtractData_Projects() []model.Projects {

	var data model.Projects
	var data_mas []model.Projects
	res, err := Db.Query("SELECT `name`,`id`, `description`, `colour`, `text_colour` FROM `projects` ;")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&data.Title, &data.Id, &data.Text, &data.BackgroundColor, &data.TextColor)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %d, %s ", data.Title, data.Id, data.Text))
		data_mas = append(data_mas, data)
	}
	fmt.Println(data_mas)
	fmt.Println(data.Title, " ", data.Id, " ", data.Text) //пример как вырывать параметры из запроса
	return data_mas

}

func ExtractDataProject_info(id string) []model.Task { //получение пользователя из бд
	var u model.Task
	var u_mas []model.Task
	res, err := Db.Query("SELECT tasks.id, tasks.name, tasks.start_date, tasks.end_date, tasks.description, projects.colour, projects.text_colour, tasks.is_finished FROM tasks INNER JOIN projects ON tasks.project_id = projects.id WHERE tasks.project_id = ? ;", id)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Id, &u.Title, &u.Start, &u.End, &u.Text, &u.BackgroundColor, &u.TextColor, &u.Is_finished)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		u_mas = append(u_mas, u)
	}
	fmt.Println(u_mas)
	fmt.Println(u.Id, " ", u.Title, " ", u.Start, " ", u.End, " ", u.BackgroundColor, " ", u.TextColor) //пример как вырывать параметры из запроса
	return u_mas

}
func ExtractDataProject(id string) model.Task { //получение пользователя из бд
	var u model.Task

	res, err := Db.Query("SELECT  id ,name, start_date, end_date,description, colour, text_colour FROM projects WHERE id = ? ;", id)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Id, &u.Title, &u.Start, &u.End, &u.Text, &u.BackgroundColor, &u.TextColor)

		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
	}
	fmt.Println(u)
	fmt.Println(u.Id, " ", u.Title, " ", u.Start, " ", u.End, " ", u.BackgroundColor, " ", u.TextColor) //пример как вырывать параметры из запроса
	return u

}
func ExtractDataTask(id string) model.Task { //получение пользователя из бд
	var u model.Task

	res, err := Db.Query("SELECT  id ,name, start_date, end_date,description, is_finished FROM tasks WHERE id = ? ;", id)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Id, &u.Title, &u.Start, &u.End, &u.Text, &u.Is_finished)
		if err != nil {
			panic(err)
		}
		//баг с Text надо проверять значение столбца
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
	}
	fmt.Println(u)
	fmt.Println(u.Id, " ", u.Title, " ", u.Start, " ", u.End, " ") //пример как вырывать параметры из запроса
	return u

}
func ExtractDataUsers() []model.User_DB { //получение пользователя из бд
	var u model.User_DB
	var users []model.User_DB
	res, err := Db.Query("SELECT `id`, `first_name`,`last_name`, `login` FROM `employees`;")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Id, &u.FName, &u.LName, &u.Login)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		users = append(users, u)
	}
	fmt.Println(users)
	fmt.Println(u.LName, " ", u.FName) //пример как вырывать параметры из запроса
	return users
}
func ExtractDataUsers_Task(id string) []model.User_DB { //получение пользователя из бд
	var u model.User_DB
	var users []model.User_DB
	res, err := Db.Query("SELECT `first_name`,`last_name`, `id` FROM `employees` where id in (select emp_id from task_for_emp where task_id = ?);", id)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.FName, &u.LName, &u.Id)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		users = append(users, u)
	}
	fmt.Println(users)
	fmt.Println(u.LName, " ", u.FName) //пример как вырывать параметры из запроса
	return users
}
func ExtractDataProject_info_ALL() []model.Task { //получение пользователя из бд
	var u model.Task
	var u_mas []model.Task
	res, err := Db.Query("SELECT tasks.id ,tasks.name,tasks.start_date,tasks.end_date, projects.colour, projects.text_colour, projects.id FROM tasks  INNER JOIN projects ON tasks.project_id = projects.id ;")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Id, &u.Title, &u.Start, &u.End, &u.BackgroundColor, &u.TextColor, &u.Project_id)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		u_mas = append(u_mas, u)
	}
	fmt.Println(u_mas)
	fmt.Println(u.Id, " ", u.Title, " ", u.Start, " ", u.End, " ", u.BackgroundColor) //пример как вырывать параметры из запроса
	return u_mas

}
func ExtractDataProject_and_Task(id string) []model.Task { //получение пользователя из бд
	var u model.Task
	var u_mas []model.Task
	fmt.Println(id)
	res, err := Db.Query("SELECT tasks.id ,tasks.name,tasks.start_date,tasks.end_date, projects.colour,projects.name,projects.id  FROM tasks  INNER JOIN projects ON tasks.project_id = projects.id WHERE tasks.id in (select task_id from task_for_emp where emp_id =?) ;", id)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Id, &u.Title, &u.Start, &u.End, &u.BackgroundColor, &u.Project_name, &u.Project_id)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		u_mas = append(u_mas, u)
	}
	fmt.Println(u_mas)
	fmt.Println(u.Id, " ", u.Title, " ", u.Start, " ", u.End, " ", u.BackgroundColor) //пример как вырывать параметры из запроса
	return u_mas

}
func AddPerson(login string, password string, first_name string, last_name string) bool {
	var count_users int

	res, err := Db.Query("SELECT COUNT(id) FROM employees WHERE login = ?;", login)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&count_users)
		if err != nil {
			panic(err)
		}
	}

	if count_users == 0 {
		result, err := Db.Exec("insert into employees (login, password, first_name,last_name, is_admin ) values (?, ?, ?,?, 0)", login, password, first_name, last_name)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.LastInsertId()) // id добавленного объекта
		fmt.Println(result.RowsAffected()) // количество затронутых строк
		return true
	} else {
		return false
	}

}

func AddProject(name string, description string, colour string, text_colour string) {
	result, err := Db.Exec("insert into projects (name, description, colour , start_date, end_date, text_colour) values (?, ?, ?, sysdate(), sysdate(), ?)", name, description, colour, text_colour)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}

func AddTask(name string, start_date string, end_date string, description string, id_project string, person []string) {
	result, err := Db.Exec("insert into tasks (name, start_date, end_date, description, project_id) values (?, STR_TO_DATE(?, '%Y-%m-%d'), STR_TO_DATE(?, '%Y-%m-%d'), ?, ?)", name, start_date, end_date, description, id_project)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк

	var id_task string

	res, err := Db.Query("SELECT MAX(`id`) FROM tasks ;")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&id_task)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(id_task)
	for i := 0; i < len(person); i++ {
		_, err1 := Db.Exec("insert into task_for_emp (task_id, emp_id) values (?,?)", id_task, person[i])
		if err1 != nil {
			panic(err)
		}
	}

	res_date, err_date := Db.Exec("update  projects set start_date = (SELECT MIN(start_date) FROM tasks WHERE project_id = ?) WHERE id = ?", id_project, id_project)
	if err_date != nil {
		panic(err_date)
	}
	fmt.Println(res_date)

	res_date1, err_date1 := Db.Exec("update  projects set end_date = (SELECT MAX(end_date) FROM tasks WHERE project_id = ?) WHERE id = ?", id_project, id_project)
	if err_date1 != nil {
		panic(err_date)
	}
	fmt.Println(res_date1)

}
func UpdateProject(id string, name string, description string, colour string, text_colour string) {
	result, err := Db.Exec("update  projects set name = ?, description = ?, colour = ?, text_colour = ? WHERE id = ?", name, description, colour, text_colour, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}
func UpdateTask(id string, name string, start_date string, end_date string, description string, person []string, is_finished bool) {
	result, err := Db.Exec("update  tasks set name=?, start_date=STR_TO_DATE(?, '%Y-%m-%d'), end_date=STR_TO_DATE(?, '%Y-%m-%d'), description =?, is_finished = ? where id = ?", name, start_date, end_date, description, is_finished, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк

	id_task := id

	res1, err1 := Db.Query("DELETE FROM task_for_emp WHERE task_id =? ;", id)
	if err1 != nil {
		panic(err)
	}
	for res1.Next() {
		err1 = res1.Scan(&id_task)
		if err1 != nil {
			panic(err)
		}
	}

	var id_project string

	res, err := Db.Query("SELECT project_id FROM tasks WHERE id = ?;", id)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&id_project)
		if err != nil {
			panic(err)
		}
	}

	_, err_date := Db.Exec("update  projects set start_date = (SELECT MIN(start_date) FROM tasks WHERE project_id = ?) WHERE id = ?", id_project, id_project)
	if err_date != nil {
		panic(err_date)
	}

	_, err_date1 := Db.Exec("update  projects set end_date = (SELECT MAX(end_date) FROM tasks WHERE project_id = ?)  WHERE id = ?", id_project, id_project)
	if err_date1 != nil {
		panic(err_date)
	}

	for i := 0; i < len(person); i++ {
		fmt.Println(i)
		_, err2 := Db.Exec("insert into task_for_emp (task_id, emp_id) values (?,?)", id_task, person[i])
		if err2 != nil {
			panic(err)
		}
		fmt.Println(i)
	}
}

func DeleteTaskFromDB(id string, task_id string) {
	result, err := Db.Exec("DELETE FROM task_for_emp WHERE task_id in (SELECT id FROM tasks WHERE id = ? AND project_id = ?)", task_id, id)
	if err != nil {
		panic(err)
	}
	result1, err1 := Db.Exec("DELETE FROM tasks WHERE id = ? AND project_id = ?", task_id, id)
	if err1 != nil {
		panic(err)
	}

	var count_tasks int

	res, err := Db.Query("SELECT COUNT(id) FROM tasks WHERE project_id = ?;", id)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&count_tasks)
		if err != nil {
			panic(err)
		}
	}

	if count_tasks > 0 {
		_, err_date := Db.Exec("update  projects set start_date = (SELECT MIN(start_date) FROM tasks WHERE project_id = ?) WHERE id = ?", id, id)
		if err_date != nil {
			panic(err_date)
		}

		_, err_date1 := Db.Exec("update  projects set end_date = (SELECT MAX(end_date) FROM tasks WHERE project_id = ?)  WHERE id = ?", id, id)
		if err_date1 != nil {
			panic(err_date)
		}
	} else {
		_, err_date := Db.Exec("update  projects set start_date = sysdate(), end_date = sysdate() WHERE id = ?", id)
		if err_date != nil {
			panic(err_date)
		}
	}
	fmt.Println(result.RowsAffected())
	fmt.Println(result1.RowsAffected()) // количество затронутых строк
}

func DeleteProjectFromDB(id string) {
	result, err := Db.Exec("DELETE FROM task_for_emp WHERE task_id in (SELECT id FROM tasks WHERE project_id = ?)", id)
	if err != nil {
		panic(err)
	}
	result1, err1 := Db.Exec("DELETE FROM tasks WHERE  project_id = ?", id)
	if err1 != nil {
		panic(err)
	}

	result2, err2 := Db.Exec("DELETE FROM projects WHERE id = ?", id)
	if err2 != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
	fmt.Println(result1.RowsAffected()) // количество затронутых строк
	fmt.Println(result2.RowsAffected())
}

func GetNotAdminUsers() []model.User_DB { //получение пользователя из бд
	var u model.User_DB
	var users []model.User_DB
	res, err := Db.Query("SELECT `id`, `first_name`,`last_name`, `login` FROM `employees` WHERE is_admin=0;")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		err = res.Scan(&u.Id, &u.FName, &u.LName, &u.Login)
		if err != nil {
			panic(err)
		}
		//fmt.Println(fmt.Sprintf("in database have %s , %s ", u.FName, u.LName))
		users = append(users, u)
	}
	fmt.Println(users)
	fmt.Println(u.LName, " ", u.FName) //пример как вырывать параметры из запроса
	return users
}
