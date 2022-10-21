package api

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/martenwallewein/go-sample-microservice/models"
	"github.com/martenwallewein/go-sample-microservice/pkg/milestones"
	"github.com/martenwallewein/go-sample-microservice/pkg/projects"
	"github.com/martenwallewein/go-sample-microservice/pkg/tasks"
)

func path(endpoint string) string {
	return fmt.Sprintf("/api/v1/%s", endpoint)
}

type RESTApiV1 struct {
	router *gin.Engine
}

func (api *RESTApiV1) Serve(addr string) error {
	return api.router.Run(addr)
}

func NewRESTApiV1() *RESTApiV1 {
	router := gin.Default()
	api := &RESTApiV1{
		router,
	}

	router.POST(path("projects/:id"), api.EditProject)
	router.DELETE(path("projects/:id"), api.DeleteProject)
	router.GET(path("projects"), api.GetProjects)
	router.PUT(path("projects"), api.AddProject)

	router.POST(path("milestones/:id"), api.EditMilestone)
	router.DELETE(path("milestones/:id"), api.DeleteMilestone)
	router.GET(path("milestones"), api.GetMilestones)
	router.PUT(path("milestones"), api.AddMilestone)

	router.POST(path("tasks/:id"), api.EditTask)
	router.DELETE(path("tasks/:id"), api.DeleteTask)
	router.GET(path("tasks"), api.GetTasks)
	router.PUT(path("tasks"), api.AddTask)

	return api
}

func (api *RESTApiV1) GetProjects(c *gin.Context) {
	projects, err := projects.GetService().GetAllProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": projects,
	})
}

func (api *RESTApiV1) EditProject(c *gin.Context) {
	id := c.Param("id")
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id, not a number"})
		return
	}
	if err := projects.GetService().EditProject(uint(idInt), project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": project.ID,
	})
}

func (api *RESTApiV1) AddProject(c *gin.Context) {

	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := projects.GetService().CreateProject(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add project"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": project.ID,
	})
}

func (api *RESTApiV1) DeleteProject(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id, not a number"})
		return
	}

	if err := projects.GetService().DeleteProject(uint(idInt)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (api *RESTApiV1) GetMilestones(c *gin.Context) {
	milestones, err := milestones.GetService().GetAllMilestones()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch milestones"})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": milestones,
	})
}

func (api *RESTApiV1) EditMilestone(c *gin.Context) {
	id := c.Param("id")
	var milestone models.Milestone
	if err := c.ShouldBindJSON(&milestone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id, not a number"})
		return
	}
	if err := milestones.GetService().EditMilestone(uint(idInt), milestone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update milestone"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": milestone.ID,
	})
}

func (api *RESTApiV1) AddMilestone(c *gin.Context) {

	var milestone models.Milestone
	if err := c.ShouldBindJSON(&milestone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := milestones.GetService().CreateMilestone(&milestone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add milestone"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": milestone.ID,
	})
}

func (api *RESTApiV1) DeleteMilestone(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id, not a number"})
		return
	}

	if err := milestones.GetService().DeleteMilestone(uint(idInt)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete milestone"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (api *RESTApiV1) GetTasks(c *gin.Context) {
	tasks, err := tasks.GetService().GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
}

func (api *RESTApiV1) EditTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id, not a number"})
		return
	}
	if err := tasks.GetService().EditTask(uint(idInt), task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": task.ID,
	})
}

func (api *RESTApiV1) AddTask(c *gin.Context) {

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := tasks.GetService().CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add task"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": task.ID,
	})
}

func (api *RESTApiV1) DeleteTask(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id, not a number"})
		return
	}

	if err := tasks.GetService().DeleteTask(uint(idInt)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
