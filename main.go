package main

import (
	"fmt"
	"todoapi/database"
	"todoapi/model"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	database.Connect()

	// task := model.Task{
	// 	Name:    "Go to shop",
	// 	Details: "Buy some shit",
	// }

	newTask := model.Task{
		Details: "Fuck some pussy",
		Done:    0,
	}

	//database.InsertNewTask(database.DB, task)
	database.UpdateTaskByID(database.DB, 37, newTask)
	database.DeleteTaskByID(database.DB, 37)

	s := database.GetAllTasks(database.DB)

	fmt.Println(s)

	defer database.DB.Close()
}
