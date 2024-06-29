package method

import (
	"Task/taskpb"
	"database/sql"
	"log"
)

func StoreNewTask(db *sql.DB, req *taskpb.Request) (*taskpb.Response, error) {
	var task taskpb.Response
	query := "INSERT INTO taskservice (title, assigned) VALUES($1, $2) RETURNING id;"
	row := db.QueryRow(query, req.Title, req.Assigned)
	if err := row.Scan(&task.Id); err != nil {
		log.Println("unable to insert task:", err)
		return nil, err
	}

	return &task, nil
}

func GetTaskById(db *sql.DB, req *taskpb.Response) (*taskpb.Task, error) {
	var task taskpb.Task
	query := "SELECT id, title, assigned FROM taskservice WHERE id = $1;"

	row := db.QueryRow(query, req.Id)
	if err := row.Scan(&task.Id, &task.Title, &task.Assigned); err != nil {
		log.Println("unable to get task:", err)
		return nil, err
	}

	return &task, nil
}

func UpdateTask(db *sql.DB, req *taskpb.Task) error {
	query := "UPDATE taskservice SET title = $1, assigned = $2 WHERE id = $3;"
	_, err := db.Exec(query, req.Title, req.Assigned, req.Id)
	if err != nil {
		log.Println("unable to update task:", err)
		return err
	}

	return nil
}

func DeleteTask(db *sql.DB, req *taskpb.Response) error {
	query := "DELETE FROM taskservice WHERE id = $1;"
	_, err := db.Exec(query, req.Id)
	if err != nil {
		log.Println("unable to delete task:", err)
		return err
	}

	return nil
}
