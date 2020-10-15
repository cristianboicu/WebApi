package database

import (
	"database/sql"
	"log"
	"todoapi/model"
)

//CreateTaskTable initialize the task table for the first time
func CreateTaskTable(db *sql.DB) {
	db.Query(`CREATE TABLE IF NOT EXISTS tasks (
        "id"	INTEGER NOT NULL UNIQUE,
        "Name"	TEXT NOT NULL,
        "Details"	TEXT,
        "Date"	TEXT,
        "Done"	INTEGER DEFAULT 0,
        PRIMARY KEY("id" AUTOINCREMENT)
)
`)
}

//GetAllTasks get all tasks from dt
func GetAllTasks(db *sql.DB) []model.Task {
	var tasks []model.Task

	row, err := db.Query("SELECT * FROM tasks ORDER BY Date")
	if err != nil {
		log.Fatal("Error while selecting all the tasks from db", err)
	}
	defer row.Close()

	for row.Next() {
		var id int
		var Name string
		var Details string
		var Date string
		var Done int
		row.Scan(&id, &Name, &Details, &Date, &Done)
		tasks = append(tasks, model.Task{ID: id, Name: Name, Details: Details, Date: Date, Done: Done})
	}

	return tasks
}

//InsertNewTask add new task in database
func InsertNewTask(db *sql.DB, task model.Task) {
	_, err := db.Exec("INSERT INTO tasks (Name, Details, Done, Date) VALUES($1, $2, 0, DATETIME('now'))",
		task.Name, task.Details)

	if err != nil {
		log.Fatal(err)
	}
}

//UpdateTaskByID update a task by id in db
func UpdateTaskByID(db *sql.DB, id int, newTask model.Task) {

	task := model.Task{}
	var ptr *int

	row, err := db.Query("SELECT * FROM tasks WHERE id = $1", id)
	if err != nil {
		log.Fatal("Error while selecting a task by id", err)
	}
	defer row.Close()

	for row.Next() {
		var (
			Name    string
			Details string
			Date    string
			Done    int
		)
		row.Scan(&id, &Name, &Details, &Date, &Done)
		task.Name = Name
		task.Details = Details
		ptr = &Done
	}

	if newTask.Name == "" {
		newTask.Name = task.Name
	}
	if newTask.Details == "" {
		newTask.Details = task.Details
	}
	if ptr == nil {
		newTask.Done = task.Done
	}

	_, err = db.Exec("UPDATE tasks SET Name = $1, Details = $2, Done = $3 WHERE id = $4",
		newTask.Name, newTask.Details, newTask.Done, id)

	if err != nil {

	}
}

//DeleteTaskByID delete a task by id :)
func DeleteTaskByID(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)

	if err != nil {
		log.Fatal("Error deleting a task", err)
	}
}
