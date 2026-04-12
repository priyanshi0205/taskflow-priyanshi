package task

import (
	"backedn_go/model"
)

func MapList(tasks []model.Task) []TaskResponse {
	response := make([]TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		response = append(response, MapSingle(task))
	}
	return response
}

func MapSingle(task model.Task) TaskResponse {
	return TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		ProjectID:   task.ProjectID,
		CreatedByID: task.CreatedByID,
		AssigneeID:  task.AssigneeID,
		DueDate:     task.DueDate,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}
