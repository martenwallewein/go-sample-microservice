package seeds

import (
	"encoding/json"
	"os"

	"github.com/martenwallewein/go-sample-microservice/models"
	"github.com/martenwallewein/go-sample-microservice/pkg/milestones"
	"github.com/martenwallewein/go-sample-microservice/pkg/projects"
	"github.com/martenwallewein/go-sample-microservice/pkg/tasks"
	"github.com/sirupsen/logrus"
)

type Seeds struct {
	Tasks      []models.Task
	Projects   []models.Project
	Milestones []models.Milestone
}

func RunSeeds(seedFilePath string) error {

	taskService := tasks.GetService()
	projectService := projects.GetService()
	milestoneService := milestones.GetService()

	dbTasks, err := taskService.GetAllTasks()
	if err != nil {
		return err
	}
	if len(dbTasks) > 0 {
		logrus.Info("Tasks already exist, skipping all migrations")
		return nil
	}

	seedFile, err := os.Open(seedFilePath)
	if err != nil {
		return err
	}
	defer seedFile.Close()
	var data Seeds
	if err := json.NewDecoder(seedFile).Decode(&data); err != nil {
		return err
	}

	for _, u := range data.Projects {
		if err := projectService.CreateProject(&u); err != nil {
			return err
		}
	}

	for _, r := range data.Milestones {
		if err := milestoneService.CreateMilestone(&r); err != nil {
			return err
		}
	}

	for _, p := range data.Tasks {
		if err := taskService.CreateTask(&p); err != nil {
			return err
		}
	}

	return nil
}
