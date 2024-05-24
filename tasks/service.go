package tasks

import (
	"awesomeProject/domain"
	"awesomeProject/repositories"
	"context"
	"github.com/gofrs/uuid"
	"golang.org/x/exp/slog"
)

type (
	TaskService interface {
		AddTask(ctx context.Context, request domain.CreateTaskRequest, userId uuid.UUID) (*domain.Task, error)
		GetTask(ctx context.Context, request domain.GetByIdRequest, userId uuid.UUID) (*domain.Task, error)
		//GetTasks
	}

	TaskServiceMiddleware func(next TaskService) TaskService

	taskService struct {
		logger         *slog.Logger
		taskRepository repositories.TaskRepository
	}
)

func NewTaskService(logger *slog.Logger, taskRepository repositories.TaskRepository) TaskService {
	return &taskService{
		logger:         logger,
		taskRepository: taskRepository,
	}
}
