package db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/haru-256/blog-ddd-uber-fx/internal/domain/models"
	"github.com/haru-256/blog-ddd-uber-fx/internal/domain/repositories"
)

type UserRepositoryImpl struct {
	users []*userData
}

func NewUserRepositoryImpl(logger *slog.Logger) *UserRepositoryImpl {
	logger.Info("Init NewUserRepositoryImpl", slog.Int("num", len(usersDB)))
	return &UserRepositoryImpl{
		users: usersDB,
	}
}

func (u *UserRepositoryImpl) GetAll(ctx context.Context) ([]*models.User, error) {
	users := make([]*models.User, len(u.users))
	for i, user := range u.users {
		tasks, err := toTaskModels(user.tasks)
		if err != nil {
			return nil, err
		}
		userId, err := models.BuildUserId(user.id)
		if err != nil {
			return nil, err
		}
		users[i], err = models.BuildUser(userId, user.name, tasks)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}

func (u *UserRepositoryImpl) GetById(ctx context.Context, id *models.UserId) (*models.User, error) {
	v, found := find(u.users, func(v *userData) bool {
		userId, err := models.BuildUserId(v.id)
		if err != nil {
			return false
		}
		return id.Equals(userId)
	})
	if !found {
		return nil, fmt.Errorf("user not found by id: %s", id)
	}
	tasks, err := toTaskModels(v.tasks)
	if err != nil {
		return nil, err
	}
	userId, err := models.BuildUserId(v.id)
	if err != nil {
		return nil, err
	}
	return models.BuildUser(userId, v.name, tasks)
}

func (u *UserRepositoryImpl) Create(ctx context.Context, user *models.User) (*models.User, error) {
	userData := &userData{
		id:   user.Id().Value(),
		name: user.Name(),
	}
	u.users = append(u.users, userData)
	return user, nil
}

func (u *UserRepositoryImpl) Update(ctx context.Context, user *models.User) (*models.User, error) {
	v, found := find(u.users, func(v *userData) bool {
		userId, err := models.BuildUserId(v.id)
		if err != nil {
			return false
		}
		return user.Id().Equals(userId)
	})
	if !found {
		return nil, fmt.Errorf("user not found by id: %s", user.Id())
	}
	v.name = user.Name()
	tasks := make([]*taskData, len(user.Tasks()))
	for i, task := range user.Tasks() {
		tasks[i] = &taskData{
			id:   task.Id().Value(),
			name: task.Name(),
		}
	}
	v.tasks = tasks
	return user, nil
}

func (u *UserRepositoryImpl) Delete(ctx context.Context, id *models.UserId) error {
	newUsers, found := deleteItem(u.users, func(v *userData) bool {
		return v.id == id.Value()
	})
	if !found {
		return fmt.Errorf("user not found by id: %s", id.Value())
	}
	u.users = newUsers
	return nil
}

type TaskRepositoryImpl struct {
	tasks []*taskData
}

func NewTaskRepositoryImpl(logger *slog.Logger) *TaskRepositoryImpl {
	logger.Info("Init NewTaskRepositoryImpl", slog.Int("num", len(tasksDB)))
	return &TaskRepositoryImpl{
		tasks: tasksDB,
	}
}

func (t *TaskRepositoryImpl) GetAll(ctx context.Context) ([]*models.Task, error) {
	tasks := make([]*models.Task, len(t.tasks))
	for i, task := range t.tasks {
		taskId, err := models.BuildTaskId(task.id)
		if err != nil {
			return nil, err
		}
		tasks[i], err = models.BuildTask(taskId, task.name)
		if err != nil {
			return nil, err
		}
	}
	return tasks, nil
}

func (t *TaskRepositoryImpl) GetById(ctx context.Context, id *models.TaskId) (*models.Task, error) {
	v, found := find(t.tasks, func(v *taskData) bool {
		return v.id == id.Value()
	})
	if !found {
		return nil, fmt.Errorf("task not found by id: %s", id.Value())
	}
	return toTaskModel(v)
}

func (t *TaskRepositoryImpl) Create(ctx context.Context, task *models.Task) (*models.Task, error) {
	taskData := &taskData{
		id:   task.Id().Value(),
		name: task.Name(),
	}
	t.tasks = append(t.tasks, taskData)
	return task, nil
}

func (t *TaskRepositoryImpl) Update(ctx context.Context, task *models.Task) (*models.Task, error) {
	v, found := find(t.tasks, func(v *taskData) bool {
		taskId, err := models.BuildTaskId(v.id)
		if err != nil {
			return false
		}
		return task.Id().Equals(taskId)
	})
	if !found {
		return nil, fmt.Errorf("task not found by id: %s", task.Id())
	}
	v.name = task.Name()
	return task, nil
}

func (t *TaskRepositoryImpl) Delete(ctx context.Context, id *models.TaskId) error {
	newTasks, found := deleteItem(t.tasks, func(v *taskData) bool {
		return v.id == id.Value()
	})
	if !found {
		return fmt.Errorf("task not found by id: %s", id.Value())
	}
	t.tasks = newTasks
	return nil
}

func find[T any](slice []T, match func(T) bool) (T, bool) {
	var zero T
	for _, v := range slice {
		if match(v) {
			return v, true
		}
	}
	return zero, false
}

func deleteItem[T any](slice []T, match func(T) bool) ([]T, bool) {
	for i, v := range slice {
		if match(v) {
			return append(slice[:i], slice[i+1:]...), true
		}
	}
	return slice, false
}

var _ repositories.UserRepository = (*UserRepositoryImpl)(nil)
var _ repositories.TaskRepository = (*TaskRepositoryImpl)(nil)

func toTaskModel(task *taskData) (*models.Task, error) {
	taskId, err := models.BuildTaskId(task.id)
	if err != nil {
		return nil, err
	}
	return models.BuildTask(taskId, task.name)
}

func toTaskModels(tasks []*taskData) ([]*models.Task, error) {
	taskModels := make([]*models.Task, len(tasks))
	var err error
	for i, task := range tasks {
		taskModels[i], err = toTaskModel(task)
		if err != nil {
			return nil, err
		}
	}
	return taskModels, nil
}
