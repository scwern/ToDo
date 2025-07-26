package service

import (
	"ToDo/internal/mocks"
	"testing"

	"ToDo/internal/domain/task"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func createTestTask() task.Task {
	t := task.NewTask("Test Task", "Description", task.StatusNew)
	t.SetID(uuid.New())
	return t
}

func TestTaskServiceGetAll(t *testing.T) {
	repo := mocks.NewTaskRepositoryInterface(t)
	service := NewTaskService(repo)

	userID := uuid.New()
	expectedTasks := []task.Task{createTestTask()}

	repo.EXPECT().GetAll(userID).Return(expectedTasks, nil).Once()

	result, err := service.GetAll(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedTasks, result)
}

func TestTaskServiceGetById(t *testing.T) {
	repo := mocks.NewTaskRepositoryInterface(t)
	service := NewTaskService(repo)

	userID := uuid.New()
	taskID := uuid.New()
	expectedTask := createTestTask()

	repo.EXPECT().GetById(userID, taskID).Return(&expectedTask, nil).Once()

	result, err := service.GetById(userID, taskID)

	assert.NoError(t, err)
	assert.Equal(t, &expectedTask, result)
}

func TestTaskServiceCreate(t *testing.T) {
	repo := mocks.NewTaskRepositoryInterface(t)
	service := NewTaskService(repo)

	newTask := task.NewTask("New Task", "Desc", task.StatusNew)
	createdTask := newTask
	createdTask.SetID(uuid.New())

	repo.EXPECT().Create(newTask).Return(createdTask, nil).Once()

	result, err := service.Create(newTask)

	assert.NoError(t, err)
	assert.Equal(t, createdTask, result)
}

func TestTaskServiceUpdate(t *testing.T) {
	repo := mocks.NewTaskRepositoryInterface(t)
	service := NewTaskService(repo)

	taskID := uuid.New()
	updatedTask := createTestTask()
	updatedTask.SetTitle("Updated Title")

	repo.EXPECT().Update(taskID, updatedTask).Return(&updatedTask, nil).Once()

	result, err := service.Update(taskID, updatedTask)

	assert.NoError(t, err)
	assert.Equal(t, &updatedTask, result)
}

func TestTaskServiceDelete(t *testing.T) {
	repo := mocks.NewTaskRepositoryInterface(t)
	service := NewTaskService(repo)

	taskID := uuid.New()

	repo.EXPECT().MarkDeleted(taskID).Return(nil).Once()

	err := service.Delete(taskID)

	assert.NoError(t, err)
}

func TestTaskServiceGetByTitle(t *testing.T) {
	repo := mocks.NewTaskRepositoryInterface(t)
	service := NewTaskService(repo)

	userID := uuid.New()
	title := "Specific Task"
	expectedTask := createTestTask()
	expectedTask.SetTitle(title)

	repo.EXPECT().GetByTitle(userID, title).Return(&expectedTask, nil).Once()

	result, err := service.GetByTitle(userID, title)

	assert.NoError(t, err)
	assert.Equal(t, &expectedTask, result)
}
