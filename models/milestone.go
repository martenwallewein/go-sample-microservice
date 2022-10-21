package models

import "gorm.io/gorm"

type Milestone struct {
	gorm.Model
	Name  string
	Tasks []Task
}
