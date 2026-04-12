package project

type createProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type updateProjectRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
