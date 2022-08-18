package dbs

import (
	"gorm.io/gorm"
)

func NewDatabaseLayer(db *gorm.DB) *DatabaseLayer {
	return &DatabaseLayer{
		db: db,
	}
}

/**
 * Project operations
**/

func (dl *DatabaseLayer) GetAllProjects() ([]Project, error) {
	var projects []Project
	result := dl.db.Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (dl *DatabaseLayer) CreateProject(project *Project) error {
	result := dl.db.Create(project)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dl *DatabaseLayer) EditProject(id uint, project Project) error {
	dbProject := Project{}
	dbProject.ID = id
	result := dl.db.Model(&dbProject).Updates(project)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dl *DatabaseLayer) DeleteProject(id uint) error {
	dbProject := Project{}
	dbProject.ID = id
	result := dl.db.Delete(&dbProject)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/**
 * Milestone operations
**/

func (dl *DatabaseLayer) GetAllMilestones() ([]Milestone, error) {
	var milestones []Milestone
	result := dl.db.Find(&milestones)
	if result.Error != nil {
		return nil, result.Error
	}
	return milestones, nil
}

func (dl *DatabaseLayer) CreateMilestone(milestone *Milestone) error {
	result := dl.db.Create(milestone)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dl *DatabaseLayer) EditMilestone(id uint, milestone Milestone) error {
	dbMilestone := Milestone{}
	dbMilestone.ID = id
	result := dl.db.Model(&dbMilestone).Updates(milestone)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dl *DatabaseLayer) DeleteMilestone(id uint) error {
	dbMilestone := Milestone{}
	dbMilestone.ID = id
	result := dl.db.Delete(&dbMilestone)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/**
 * Task operations
**/

func (dl *DatabaseLayer) GetAllTasks() ([]Task, error) {
	var tasks []Task
	result := dl.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (dl *DatabaseLayer) CreateTask(task *Task) error {
	result := dl.db.Create(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dl *DatabaseLayer) EditTask(id uint, task Task) error {
	dbTask := Task{}
	dbTask.ID = id
	result := dl.db.Model(&dbTask).Updates(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dl *DatabaseLayer) DeleteTask(id uint) error {
	dbTask := Task{}
	dbTask.ID = id
	result := dl.db.Delete(&dbTask)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
