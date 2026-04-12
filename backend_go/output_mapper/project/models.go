package project

import (
	taskoutput "backedn_go/output_mapper/task"
	"time"
)

type projectResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID     string    `json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type projectDetailResponse struct {
	projectResponse
	Tasks []taskoutput.TaskResponse `json:"tasks"`
}
