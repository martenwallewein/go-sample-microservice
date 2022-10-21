package projects

import (
	"sync"

	"github.com/martenwallewein/go-sample-microservice/dbs"
	"github.com/martenwallewein/go-sample-microservice/models"
	"gorm.io/gorm"
)

type ProjectService struct {
	db *gorm.DB
}

var projectService *ProjectService
var initOnce sync.Once

func GetService() *ProjectService {
	initOnce.Do(initService)
	return projectService
}

func initService() {
	db := dbs.GetDB()
	projectService = NewProjectService(db)
}

func NewProjectService(db *gorm.DB) *ProjectService {
	return &ProjectService{
		db: db,
	}
}

/**
 * Project operations
**/

func (dl *ProjectService) GetAllProjects() ([]models.Project, error) {
	var projects []models.Project
	result := dl.db.Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (dl *ProjectService) CreateProject(project *models.Project) error {
	result := dl.db.Create(project)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dl *ProjectService) EditProject(id uint, project models.Project) error {
	dbProject := models.Project{}
	dbProject.ID = id
	result := dl.db.Model(&dbProject).Updates(project)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dl *ProjectService) DeleteProject(id uint) error {
	dbProject := models.Project{}
	dbProject.ID = id
	result := dl.db.Delete(&dbProject)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
