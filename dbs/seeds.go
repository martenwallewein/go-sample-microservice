package dbs

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

type Seeds struct {
	Tasks      []Task
	Projects   []Project
	Milestones []Milestone
}

func RunSeeds(dl *DatabaseLayer, seedFilePath string) error {

	dbTasks, err := dl.GetAllTasks()
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
		if err := dl.CreateProject(&u); err != nil {
			return err
		}
	}

	for _, r := range data.Milestones {
		if err := dl.CreateMilestone(&r); err != nil {
			return err
		}
	}

	for _, p := range data.Tasks {
		if err := dl.CreateTask(&p); err != nil {
			return err
		}
	}

	return nil
}
