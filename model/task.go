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

func AddTask(name string) (*Task, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	task := Task{
		ID:         id,
		Name:       name,
		IsFinished: false,
	}

	err = db.Create(&task).Error

	return &task, err
}
