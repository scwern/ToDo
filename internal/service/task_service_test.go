package service

import (
	"ToDo/internal/domain/task"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) GetAll(userID uuid.UUID) ([]task.Task, error) {
	args := m.Called(userID)
	return args.Get(0).([]task.Task), args.Error(1)
}

func (m *MockTaskRepository) GetById(userID, id uuid.UUID) (*task.Task, error) {
	args := m.Called(userID, id)
	return args.Get(0).(*task.Task), args.Error(1)
}

func (m *MockTaskRepository) Create(t task.Task) (task.Task, error) {
	args := m.Called(t)
	return args.Get(0).(task.Task), args.Error(1)
}

func (m *MockTaskRepository) Update(id uuid.UUID, updated task.Task) (*task.Task, error) {
	args := m.Called(id, updated)
	return args.Get(0).(*task.Task), args.Error(1)
}

func (m *MockTaskRepository) MarkDeleted(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskRepository) GetByTitle(userID uuid.UUID, title string) (*task.Task, error) {
	args := m.Called(userID, title)
	return args.Get(0).(*task.Task), args.Error(1)
}

func createTestTask() task.Task {
	t := task.NewTask("Test", "Description", task.StatusNew)
	t.SetID(uuid.New())
	return t
}

func TestTaskService_GetAll(t *testing.T) {
	repo := new(MockTaskRepository)
	service := NewTaskService(repo)
	userID := uuid.New()

	task1 := createTestTask()
	task2 := createTestTask()
	expectedTasks := []task.Task{task1, task2}

	repo.On("GetAll", userID).Return(expectedTasks, nil)
	tasks, err := service.GetAll(userID)
	assert.NoError(t, err)
	assert.Equal(t, expectedTasks, tasks)
}

func TestTaskService_GetAll_Error(t *testing.T) {
	repo := new(MockTaskRepository)
	service := NewTaskService(repo)
	userID := uuid.New()
	repo.On("GetAll", userID).Return([]task.Task{}, errors.New("error"))
	_, err := service.GetAll(userID)
	assert.Error(t, err)
}

func TestTaskService_GetById(t *testing.T) {
	repo := new(MockTaskRepository)
	service := NewTaskService(repo)
	userID := uuid.New()
	taskID := uuid.New()

	expectedTask := createTestTask()
	expectedTask.SetID(taskID)

	repo.On("GetById", userID, taskID).Return(&expectedTask, nil)
	result, err := service.GetById(userID, taskID)
	assert.NoError(t, err)
	assert.Equal(t, &expectedTask, result)
}

func TestTaskService_Create(t *testing.T) {
	repo := new(MockTaskRepository)
	service := NewTaskService(repo)

	newTask := task.NewTask("New", "Desc", task.StatusNew)
	createdTask := newTask
	createdTask.SetID(uuid.New())

	repo.On("Create", newTask).Return(createdTask, nil)
	result, err := service.Create(newTask)
	assert.NoError(t, err)
	assert.Equal(t, createdTask, result)
}

func TestTaskService_Update(t *testing.T) {
	repo := new(MockTaskRepository)
	service := NewTaskService(repo)
	taskID := uuid.New()

	updatedTask := createTestTask()
	updatedTask.SetTitle("Updated")
	updatedTask.SetID(taskID)

	repo.On("Update", taskID, updatedTask).Return(&updatedTask, nil)
	result, err := service.Update(taskID, updatedTask)
	assert.NoError(t, err)
	assert.Equal(t, &updatedTask, result)
}

func TestTaskService_Delete(t *testing.T) {
	repo := new(MockTaskRepository)
	service := NewTaskService(repo)
	taskID := uuid.New()
	repo.On("MarkDeleted", taskID).Return(nil)
	err := service.Delete(taskID)
	assert.NoError(t, err)
}
