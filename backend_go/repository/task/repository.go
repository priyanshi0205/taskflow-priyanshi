package task

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

func (r *repository) ListByProject(projectID, status, assigneeID string) ([]model.Task, error) {
	query := r.db.Where("project_id = ?", projectID)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if assigneeID != "" {
		query = query.Where("assignee_id = ?", assigneeID)
	}

	var tasks []model.Task
	if err := query.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *repository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *repository) FindByID(id string) (*model.Task, error) {
	var task model.Task
	if err := r.db.First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *repository) IsOwnedByUser(taskID, userID string) (bool, error) {
	var count int64
	err := r.db.Table("tasks").
		Joins("JOIN projects ON projects.id = tasks.project_id").
		Where("tasks.id = ? AND projects.owner_id = ?", taskID, userID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repository) IsCreator(taskID, userID string) (bool, error) {
	var count int64
	err := r.db.Table("tasks").
		Where("id = ? AND created_by_id = ?", taskID, userID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repository) Update(task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&model.Task{}, "id = ?", id).Error
}
