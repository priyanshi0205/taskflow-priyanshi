package project

import (
	"backedn_go/model"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) ListByUser(userID string) ([]model.Project, error) {
	var projects []model.Project
	query := `
		SELECT DISTINCT p.id, p.name, p.description, p.owner_id, p.created_at
		FROM projects p
		LEFT JOIN tasks t ON p.id = t.project_id
		WHERE p.owner_id = ? OR t.assignee_id = ?
	`
	if err := r.db.Raw(query, userID, userID).Scan(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *repository) Create(name, description, ownerID string) (*model.Project, error) {
	project := &model.Project{
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
	}
	if err := r.db.Create(project).Error; err != nil {
		return nil, err
	}
	return project, nil
}

func (r *repository) FindByID(id string) (*model.Project, error) {
	var project model.Project
	if err := r.db.First(&project, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *repository) FindByIDWithTasks(id string) (*model.Project, error) {
	var project model.Project
	if err := r.db.Preload("Tasks").First(&project, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

type taskStatusCountRow struct {
	Status string
	Count  int64
}

func (r *repository) CountByStatus(projectID string) (map[string]int64, error) {
	var rows []taskStatusCountRow
	err := r.db.Table("tasks").
		Select("status, COUNT(*) as count").
		Where("project_id = ?", projectID).
		Group("status").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	result := map[string]int64{
		"todo":        0,
		"in_progress": 0,
		"done":        0,
	}
	for _, row := range rows {
		result[row.Status] = row.Count
	}

	return result, nil
}

type taskAssigneeCountRow struct {
	AssigneeID string
	Count      int64
}

func (r *repository) CountByAssignee(projectID string) (map[string]int64, error) {
	var rows []taskAssigneeCountRow
	err := r.db.Table("tasks").
		Select("COALESCE(assignee_id::text, 'unassigned') as assignee_id, COUNT(*) as count").
		Where("project_id = ?", projectID).
		Group("COALESCE(assignee_id::text, 'unassigned')").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	result := make(map[string]int64, len(rows))
	for _, row := range rows {
		result[row.AssigneeID] = row.Count
	}

	return result, nil
}

func (r *repository) Update(project *model.Project) error {
	return r.db.Save(project).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&model.Project{}, "id = ?", id).Error
}
