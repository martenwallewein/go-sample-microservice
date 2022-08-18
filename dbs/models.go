package dbs

import (
	"gorm.io/gorm"
)

type DatabaseLayer struct {
	db *gorm.DB
}

type Task struct {
	gorm.Model
	Name        string
	Description string
	Done        bool
}
type Milestone struct {
	gorm.Model
	Name  string
	Tasks []Task
}

type Project struct {
	gorm.Model
	Name        string
	Description string
	Milestones  []Milestone
}
