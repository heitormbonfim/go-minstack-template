package task_repositories

import (
	task_entities "go-minstack-task/internal/tasks/entities"

	"github.com/go-minstack/repository"
	"gorm.io/gorm"
)

type TaskRepository struct {
	*repository.Repository[task_entities.Task]
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{repository.NewRepository[task_entities.Task](db)}
}

func (r *TaskRepository) FindByUserID(userID uint) ([]task_entities.Task, error) {
	return r.FindAll(repository.Where("user_id = ?", userID))
}

func (r *TaskRepository) FindByIDAndUserID(id, userID uint) (*task_entities.Task, error) {
	return r.FindOne(repository.Where("id = ? AND user_id = ?", id, userID))
}
