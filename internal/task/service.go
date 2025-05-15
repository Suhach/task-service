package task

import "context"

type Service interface {
	GetTask(ctx context.Context, id int) (*Task, error)
	ListTasks(ctx context.Context) ([]*Task, error)
	GetTaskByUserID(ctx context.Context, userID int) ([]*Task, error)
	CreateTask(ctx context.Context, task *Task) error
	UpdateTask(ctx context.Context, id int, task *Task) error
	DeleteTask(ctx context.Context, id int) error
}

type TaskService struct {
	repo *TaskREPO
}

func NewTaskService(repo *TaskREPO) *TaskService {
	return &TaskService{repo}
}

func (s *TaskService) GetTask(ctx context.Context, id int) (*Task, error) {
	return s.repo.GetTask(ctx, id)
}

func (s *TaskService) ListTasks(ctx context.Context) ([]*Task, error) {
	return s.repo.ListTasks(ctx)
}

func (s *TaskService) GetTaskByUserID(ctx context.Context, userID int) ([]*Task, error) {
	return s.repo.GetTaskByUserID(ctx, userID)
}

func (s *TaskService) CreateTask(ctx context.Context, task *Task) error {
	return s.repo.CreateTask(ctx, task)
}

func (s *TaskService) UpdateTask(ctx context.Context, id int, task *Task) error {
	return s.repo.UpdateTask(ctx, id, task)
}

func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	return s.repo.DeleteTask(ctx, id)
}
