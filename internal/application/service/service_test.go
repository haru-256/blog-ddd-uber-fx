package service_test

import (
	"context"
	"testing"

	"github.com/haru-256/blog-ddd-uber-fx/internal/application/dto"
	"github.com/haru-256/blog-ddd-uber-fx/internal/application/service"
	"github.com/haru-256/blog-ddd-uber-fx/internal/domain/models"
	mock_repository "github.com/haru-256/blog-ddd-uber-fx/internal/mocks/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserServiceImpl(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUserRepository(ctrl)
	svc := service.NewUserServiceImpl(mockRepo)

	t.Run("Create", func(t *testing.T) {
		name := "test-user"
		expectedUser := mustNewUser(name, []*models.Task{})

		mockRepo.EXPECT().Create(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, user *models.User) (*models.User, error) {
			assert.Equal(t, user.Name(), name)
			return expectedUser, nil
		})

		got, err := svc.Create(ctx, name)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser.Id().Value(), got.Id)
		assert.Equal(t, expectedUser.Name(), got.Name)
	})

	t.Run("GetAll", func(t *testing.T) {
		users := []*models.User{
			mustNewUser("u1", []*models.Task{}),
			mustNewUser("u2", []*models.Task{}),
		}
		mockRepo.EXPECT().GetAll(ctx).Return(users, nil)

		got, err := svc.GetAll(ctx)
		assert.NoError(t, err)
		assert.Len(t, got, 2)
	})

	t.Run("Update", func(t *testing.T) {
		name := "new-name"
		existingUser := mustNewUser("old-name", []*models.Task{})
		updatedUser := mustNewUser(name, []*models.Task{})

		taskDTOs := []*dto.TaskDTO{
			{Id: models.NewTaskId().Value(), Name: "task1"},
		}

		mockRepo.EXPECT().GetById(ctx, gomock.Any()).Return(existingUser, nil)
		mockRepo.EXPECT().Update(ctx, gomock.Any()).Return(updatedUser, nil)

		got, err := svc.Update(ctx, existingUser.Id().Value(), name, taskDTOs)
		assert.NoError(t, err)
		assert.Equal(t, name, got.Name)
	})
}

func TestTaskServiceImpl(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockTaskRepository(ctrl)
	svc := service.NewTaskServiceImpl(mockRepo)

	t.Run("Create", func(t *testing.T) {
		name := "test-task"
		expectedTask := mustNewTask(name)

		mockRepo.EXPECT().Create(ctx, gomock.Any()).Return(expectedTask, nil)

		got, err := svc.Create(ctx, name)
		assert.NoError(t, err)
		assert.Equal(t, expectedTask.Id().Value(), got.Id)
	})

	t.Run("GetAll", func(t *testing.T) {
		tasks := []*models.Task{
			mustNewTask("t1"),
		}
		mockRepo.EXPECT().GetAll(ctx).Return(tasks, nil)

		got, err := svc.GetAll(ctx)
		assert.NoError(t, err)
		assert.Len(t, got, 1)
	})
}

func mustNewUser(name string, tasks []*models.Task) *models.User {
	u, err := models.NewUser(name, tasks)
	if err != nil {
		panic(err)
	}
	return u
}

func mustNewTask(name string) *models.Task {
	t, err := models.NewTask(name)
	if err != nil {
		panic(err)
	}
	return t
}
