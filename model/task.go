package model

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Task struct {
	ID         uuid.UUID
	Name       string
	IsFinished bool
}

func GetTasks() ([]Task, error) {
	var tasks []Task
	err := db.Find(&tasks).Error
	return tasks, err
}
