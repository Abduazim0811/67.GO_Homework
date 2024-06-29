package taskservice

import (
	"Task/internal/method"
	"Task/taskpb"
	"context"
	"database/sql"
	"log"
)

type TaskService struct {
	taskpb.UnimplementedTaskserviceServer
	db *sql.DB
}

func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{db: db}
}

func (t *TaskService) CreateTask(ctx context.Context, req *taskpb.Request) (*taskpb.Response, error) {
	task, err := method.StoreNewTask(t.db, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &taskpb.Response{Id: task.Id}, nil
}

func (t *TaskService) GetTask(ctx context.Context, req *taskpb.Response) (*taskpb.Task, error) {
	task, err := method.GetTaskById(t.db, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &taskpb.Task{
		Id:       task.Id,
		Title:    task.Title,
		Assigned: task.Assigned,
	}, nil
}

func (t *TaskService) UpdateTask(ctx context.Context, req *taskpb.Task) (*taskpb.TaskRequest, error) {
	err := method.UpdateTask(t.db, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &taskpb.TaskRequest{Message: "Updated task"}, nil
}

func (t *TaskService) DeleteTask(ctx context.Context, req *taskpb.Response) (*taskpb.TaskRequest, error) {
	err := method.DeleteTask(t.db, req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &taskpb.TaskRequest{Message: "Deleted task"}, nil
}
