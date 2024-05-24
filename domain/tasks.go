package domain

import (
	"github.com/gofrs/uuid"
	"time"
)

type (
	Task struct {
		ID          uuid.UUID `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Status      int       `json:"status"`
		IsFavourite bool      `json:"isFavourite"`
		CreatedAT   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		UserID      uuid.UUID `json:"userID"`
	}

	TaskResponse struct {
		Task *Task `json:"task"`
	}

	TasksResponse struct {
		Tasks []Task `json:"tasks"`
	}

	CreateTaskRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	UpdateTaskRequest struct {
		ID          uuid.UUID `json:"id"`
		Title       string    `json:"title,omitempty"`
		Description string    `json:"description,omitempty"`
		Status      int       `json:"status,omitempty"`
	}

	FavouriteRequest struct {
		ID uuid.UUID `json:"id" validate:"required"`
	}

	UnFavouriteRequest struct {
		ID uuid.UUID `json:"id" validate:"required"`
	}

	GetByIdRequest struct {
		ID uuid.UUID `json:"id" validate:"required"`
	}

	GetTasksRequest struct {
		UserID uuid.UUID `json:"userID" validate:"required"`
	}

	GetListOfFavouritesRequest struct {
		UserID uuid.UUID `json:"userID" validate:"required"`
	}
)
