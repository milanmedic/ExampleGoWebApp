package taskrepo

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	Task "taskapi.com/m/v2/src/models/task"
	"taskapi.com/m/v2/src/persistence/db"
)

type taskRepo struct {
	db *db.Database
}

type TaskRepositer interface {
	GetTask(id string) (Task.Tasker, error)
	GetAllTasks() ([]Task.Tasker, error)
	AddTask(taskInput Task.NewTaskInputDto) error
	DeleteTask(id string) error
	EditTask(id string, newTaskInfo Task.EditTaskDto) error
	CompleteTask(id string) error
}

func CreateTaskRepository(db *db.Database) TaskRepositer {
	tr := &taskRepo{db: db}
	tr.SetupTaskTable()
	return tr
}

//TODO: Encapsulate in SETUP DATABASE for example, and create all the tables there
func (tr taskRepo) SetupTaskTable() {
	db := tr.db.GetDbConnection()
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS tasks(
		id uuid PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL,
		description TEXT,
		status BOOLEAN DEFAULT FALSE
	);`)
	if err != nil {
		panic(err)
	}
}

func (tr taskRepo) GetTask(taskId string) (Task.Tasker, error) {
	statement := `SELECT * FROM tasks
	where id = $1`
	db := tr.db.GetDbConnection()
	var id uuid.UUID
	var name string
	var description string
	var status bool
	var task *Task.TaskDto

	row := db.QueryRow(statement, taskId)

	switch err := row.Scan(&id, &name, &description, &status); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return nil, nil
	case nil:
		task = &Task.TaskDto{Id: id, Name: name, Description: description, Status: status}
	default:
		panic(err)
	}
	return task, nil
}

func (tr taskRepo) GetAllTasks() ([]Task.Tasker, error) {
	statement := `SELECT * FROM tasks`
	db := tr.db.GetDbConnection()

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []Task.Tasker{}

	for rows.Next() {
		task := &Task.TaskDto{}

		err = rows.Scan(&task.Id, &task.Name, &task.Description, &task.Status)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tr taskRepo) AddTask(taskInput Task.NewTaskInputDto) error {
	task := Task.CreateTask(taskInput.Name, taskInput.Description)
	statement := `INSERT INTO tasks (id, name, description, status)
	VALUES ($1, $2, $3, $4)`
	db := tr.db.GetDbConnection()

	_, err := db.Exec(statement, task.Id, task.Name, task.Description, task.Status)
	if err != nil {
		return err
	}

	return nil
}

func (tr taskRepo) DeleteTask(id string) error {
	statement := `DELETE FROM tasks WHERE id = $1`

	db := tr.db.GetDbConnection()

	_, err := db.Exec(statement, id)

	if err != nil {
		return err
	}

	return nil
}

func (tr taskRepo) EditTask(id string, newTaskInfo Task.EditTaskDto) error {
	statement := `
		UPDATE tasks
		SET name = $1,
			description = $2,
			status = $3
		WHERE id = $4`

	db := tr.db.GetDbConnection()

	res, err := db.Exec(statement, newTaskInfo.Name, newTaskInfo.Description, newTaskInfo.Status, id)

	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (tr taskRepo) CompleteTask(id string) error {
	statement := `UPDATE tasks
				SET status = $1
				WHERE id = $2`
	db := tr.db.GetDbConnection()

	res, err := db.Exec(statement, true, id)

	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
