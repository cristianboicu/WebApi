package database

//CreateTaskTable initialize the task table for the first time
func CreateTaskTable() {
	DB.Query(`CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    name TEXT,
    details TEXT,
    date DATE,
    done VARCHAR(4)
)
`)
}
