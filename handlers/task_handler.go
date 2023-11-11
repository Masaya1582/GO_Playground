package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"GO_Playground/models"
)

var tasks []models.Task

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func AddTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, newTask)
}

func UpdateTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Done = !tasks[i].Done
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
