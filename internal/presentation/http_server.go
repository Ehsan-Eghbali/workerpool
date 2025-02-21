package presentation

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"workerpool/internal/app"
)

type HTTPServer struct {
	taskService *app.TaskService
}

func NewHTTPServer(taskService *app.TaskService) *HTTPServer {
	return &HTTPServer{taskService: taskService}
}

func (s *HTTPServer) Start() {
	r := gin.Default()

	r.POST("/task/:id", s.handleTask)

	r.Run(":8080")
}

func (s *HTTPServer) handleTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req struct {
		Data string `json:"data"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	s.taskService.ProcessTask(id, req.Data)
	c.JSON(http.StatusOK, gin.H{"status": "Task received"})
}
