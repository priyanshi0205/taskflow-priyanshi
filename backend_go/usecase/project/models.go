package project

type CreateProjectInput struct {
	Name        string
	Description string
}

type UpdateProjectInput struct {
	Name        *string
	Description *string
}

type ProjectStatsOutput struct {
	ByStatus   map[string]int64 `json:"by_status"`
	ByAssignee map[string]int64 `json:"by_assignee"`
}
