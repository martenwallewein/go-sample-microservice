package milestones

import (
	"sync"

	"github.com/martenwallewein/go-sample-microservice/dbs"
	"github.com/martenwallewein/go-sample-microservice/models"
	"gorm.io/gorm"
)

type MilestoneService struct {
	db *gorm.DB
}

var milestoneService *MilestoneService
var initOnce sync.Once

func GetService() *MilestoneService {
	initOnce.Do(initService)
	return milestoneService
}

func initService() {
	db := dbs.GetDB()
	milestoneService = NewMilestoneService(db)
}

func NewMilestoneService(db *gorm.DB) *MilestoneService {
	return &MilestoneService{
		db: db,
	}
}

/**
 * Milestone operations
**/

func (dl *MilestoneService) GetAllMilestones() ([]models.Milestone, error) {
	var milestones []models.Milestone
	result := dl.db.Find(&milestones)
	if result.Error != nil {
		return nil, result.Error
	}
	return milestones, nil
}

func (dl *MilestoneService) CreateMilestone(milestone *models.Milestone) error {
	result := dl.db.Create(milestone)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dl *MilestoneService) EditMilestone(id uint, milestone models.Milestone) error {
	dbMilestone := models.Milestone{}
	dbMilestone.ID = id
	result := dl.db.Model(&dbMilestone).Updates(milestone)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dl *MilestoneService) DeleteMilestone(id uint) error {
	dbMilestone := models.Milestone{}
	dbMilestone.ID = id
	result := dl.db.Delete(&dbMilestone)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
