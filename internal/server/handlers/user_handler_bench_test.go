package handlers_test

import (
	userDomain "ToDo/internal/domain/user"
	userDto "ToDo/internal/dto/user"
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

func BenchmarkUserHandlerGetAll(b *testing.B) {
	mockRepo := new(mocks.UserRepositoryInterface)

	mockUsers := make([]userDomain.User, 100)
	for i := 0; i < 100; i++ {
		u := userDomain.NewUser("User "+strconv.Itoa(i), "user"+strconv.Itoa(i)+"@example.com", "password")
		u.SetID(uuid.New())
		mockUsers[i] = u
	}

	mockRepo.On("GetAll").Return(mockUsers, nil)

	userService := service.NewUserService(mockRepo)
	handler := handlers.NewUserHandler(userService)

	router := gin.New()
	router.GET("/users", handler.GetAll)

	req, _ := http.NewRequest("GET", "/users", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}

func BenchmarkUserHandlerCreate(b *testing.B) {
	mockRepo := new(mocks.UserRepositoryInterface)

	newUser := userDomain.NewUser("Test User", "test@example.com", "password")
	newUser.SetID(uuid.New())

	mockRepo.On("Create", mock.AnythingOfType("user.User")).Return(newUser, nil)
	mockRepo.On("GetByEmail", "test@example.com").Return(nil, nil)

	userService := service.NewUserService(mockRepo)
	handler := handlers.NewUserHandler(userService)

	router := gin.New()
	router.POST("/users", handler.Create)

	userDTO := userDto.CreateUserDTO{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password",
	}
	body, _ := json.Marshal(userDTO)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
	}
}
