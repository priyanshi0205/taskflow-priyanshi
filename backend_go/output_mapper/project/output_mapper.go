package project

import (
	"backedn_go/model"
	taskoutput "backedn_go/output_mapper/task"
)

func MapList(projects []model.Project) []projectResponse {
	response := make([]projectResponse, 0, len(projects))
	for _, project := range projects {
		response = append(response, mapSingle(project))
	}
	return response
}

func MapSingle(project model.Project) projectResponse {
	return mapSingle(project)
}

func MapWithTasks(project model.Project) projectDetailResponse {
	tasks := make([]taskoutput.TaskResponse, 0, len(project.Tasks))
	for _, task := range project.Tasks {
		tasks = append(tasks, taskoutput.MapSingle(task))
	}

	return projectDetailResponse{
		projectResponse: mapSingle(project),
		Tasks:           tasks,
	}
}

func mapSingle(project model.Project) projectResponse {
	return projectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		OwnerID:     project.OwnerID,
		CreatedAt:   project.CreatedAt,
	}
}
