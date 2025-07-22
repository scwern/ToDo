package handlers_test

import (
	taskDomain "ToDo/internal/domain/task"
	taskDto "ToDo/internal/dto/task"
	"ToDo/internal/mocks"
	"ToDo/internal/server/handlers"
	"ToDo/internal/service"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func BenchmarkTaskHandlerGetAll(b *testing.B) {
	mockRepo := new(mocks.TaskRepositoryInterface)
	userID := uuid.New()

	mockTasks := make([]taskDomain.Task, 100)
	for i := 0; i < 100; i++ {
		t := taskDomain.NewTask("Task "+strconv.Itoa(i), "Description", taskDomain.StatusNew)
		t.SetID(uuid.New())
		t.SetUserID(userID)
		mockTasks[i] = t
	}

	mockRepo.On("GetAll", userID).Return(mockTasks, nil)

	taskService := service.NewTaskService(mockRepo)
	handler := handlers.NewTaskHandler(taskService)

	router := gin.New()
	router.GET("/tasks", handler.GetAll)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	req.AddCookie(&http.Cookie{Name: "user_id", Value: userID.String()})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}

func BenchmarkTaskHandlerCreate(b *testing.B) {
	mockRepo := new(mocks.TaskRepositoryInterface)
	userID := uuid.New()

	newTask := taskDomain.NewTask("Test Task", "Description", taskDomain.StatusNew)
	newTask.SetID(uuid.New())
	newTask.SetUserID(userID)

	mockRepo.On("Create", mock.AnythingOfType("task.Task")).Return(newTask, nil)
	mockRepo.On("GetByTitle", userID, "Test Task").Return(nil, nil)

	taskService := service.NewTaskService(mockRepo)
	handler := handlers.NewTaskHandler(taskService)

	router := gin.New()
	router.POST("/tasks", handler.Create)

	taskDTO := taskDto.CreateTaskDTO{
		Title:       "Test Task",
		Description: "Description",
		Status:      int(taskDomain.StatusNew),
	}
	body, _ := json.Marshal(taskDTO)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
		req.AddCookie(&http.Cookie{Name: "user_id", Value: userID.String()})
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}
