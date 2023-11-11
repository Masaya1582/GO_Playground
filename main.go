package main

import (
	"github.com/gin-gonic/gin"
	"GO_Playground/handlers"
)

func main() {
	r := gin.Default()

	r.GET("/tasks", handlers.GetTasks)
	r.POST("/tasks", handlers.AddTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)

	r.Run(":8080")
}
