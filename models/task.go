package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name        string
	Description string
	Done        bool
}
