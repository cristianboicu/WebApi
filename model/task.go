package model

//Task model structure...
type Task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Details string `json:"details"`
	Date    string `json:"date"`
	Done    int    `json:"done"`
}
