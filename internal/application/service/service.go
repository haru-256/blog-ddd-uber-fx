package service

import (
	"context"

	"github.com/haru-256/blog-ddd-uber-fx/internal/application/dto"
	"github.com/haru-256/blog-ddd-uber-fx/internal/domain/models"
	"github.com/haru-256/blog-ddd-uber-fx/internal/domain/repositories"
)

//go:generate go tool mockgen -source=$GOFILE -destination=../../mocks/service/service_mock.go -package=mock_service
type UserService interface {
	Create(ctx context.Context, name string) (*dto.UserDTO, error)
	GetAll(ctx context.Context) ([]*dto.UserDTO, error)
	GetById(ctx context.Context, id string) (*dto.UserDTO, error)
	Update(ctx context.Context, id string, name string, tasks []*dto.TaskDTO) (*dto.UserDTO, error)
	Delete(ctx context.Context, id string) error
}

//go:generate go tool mockgen -source=$GOFILE -destination=../../mocks/service/service_mock.go -package=mock_service
type TaskService interface {
	Create(ctx context.Context, name string) (*dto.TaskDTO, error)
	GetAll(ctx context.Context) ([]*dto.TaskDTO, error)
	GetById(ctx context.Context, id string) (*dto.TaskDTO, error)
	Update(ctx context.Context, id string, name string) (*dto.TaskDTO, error)
	Delete(ctx context.Context, id string) error
}

type UserServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewUserServiceImpl(userRepo repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) Create(ctx context.Context, name string) (*dto.UserDTO, error) {
	user, err := models.NewUser(name, []*models.Task{})
	if err != nil {
		return nil, err
	}
	createdUser, err := s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return toUserDTO(createdUser), nil
}

func (s *UserServiceImpl) GetAll(ctx context.Context) ([]*dto.UserDTO, error) {
	users, err := s.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	userDTOs := make([]*dto.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = toUserDTO(user)
	}
	return userDTOs, nil
}

func (s *UserServiceImpl) GetById(ctx context.Context, id string) (*dto.UserDTO, error) {
	userId, err := models.BuildUserId(id)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return toUserDTO(user), nil
}

func (s *UserServiceImpl) Update(ctx context.Context, id string, name string, tasks []*dto.TaskDTO) (*dto.UserDTO, error) {
	userId, err := models.BuildUserId(id)
	if err != nil {
		return nil, err
	}
	user, err := s.userRepo.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}
	if err = user.SetName(name); err != nil {
		return nil, err
	}
	taskModels := make([]*models.Task, len(tasks))
	for i, task := range tasks {
		taskModels[i], err = toTaskModel(task)
		if err != nil {
			return nil, err
		}
	}
	user.SetTasks(taskModels)
	updatedUser, err := s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return toUserDTO(updatedUser), nil
}

func (s *UserServiceImpl) Delete(ctx context.Context, id string) error {
	userId, err := models.BuildUserId(id)
	if err != nil {
		return err
	}
	return s.userRepo.Delete(ctx, userId)
}

func toUserDTO(user *models.User) *dto.UserDTO {
	taskDTOs := make([]*dto.TaskDTO, len(user.Tasks()))
	for i, task := range user.Tasks() {
		taskDTOs[i] = toTaskDTO(task)
	}
	return &dto.UserDTO{
		Id:    user.Id().Value(),
		Name:  user.Name(),
		Tasks: taskDTOs,
	}
}

type TaskServiceImpl struct {
	taskRepo repositories.TaskRepository
}

func NewTaskServiceImpl(taskRepo repositories.TaskRepository) *TaskServiceImpl {
	return &TaskServiceImpl{
		taskRepo: taskRepo,
	}
}

func (s *TaskServiceImpl) Create(ctx context.Context, name string) (*dto.TaskDTO, error) {
	task, err := models.NewTask(name)
	if err != nil {
		return nil, err
	}
	createdTask, err := s.taskRepo.Create(ctx, task)
	if err != nil {
		return nil, err
	}
	return toTaskDTO(createdTask), nil
}

func (s *TaskServiceImpl) GetAll(ctx context.Context) ([]*dto.TaskDTO, error) {
	tasks, err := s.taskRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	taskDTOs := make([]*dto.TaskDTO, len(tasks))
	for i, task := range tasks {
		taskDTOs[i] = toTaskDTO(task)
	}
	return taskDTOs, nil
}

func (s *TaskServiceImpl) GetById(ctx context.Context, id string) (*dto.TaskDTO, error) {
	taskId, err := models.BuildTaskId(id)
	if err != nil {
		return nil, err
	}
	task, err := s.taskRepo.GetById(ctx, taskId)
	if err != nil {
		return nil, err
	}
	return toTaskDTO(task), nil
}

func (s *TaskServiceImpl) Update(ctx context.Context, id string, name string) (*dto.TaskDTO, error) {
	taskId, err := models.BuildTaskId(id)
	if err != nil {
		return nil, err
	}
	task, err := s.taskRepo.GetById(ctx, taskId)
	if err != nil {
		return nil, err
	}
	if err = task.SetName(name); err != nil {
		return nil, err
	}
	updatedTask, err := s.taskRepo.Update(ctx, task)
	if err != nil {
		return nil, err
	}
	return toTaskDTO(updatedTask), nil
}

func (s *TaskServiceImpl) Delete(ctx context.Context, id string) error {
	taskId, err := models.BuildTaskId(id)
	if err != nil {
		return err
	}
	return s.taskRepo.Delete(ctx, taskId)
}

func toTaskDTO(task *models.Task) *dto.TaskDTO {
	return &dto.TaskDTO{
		Id:   task.Id().Value(),
		Name: task.Name(),
	}
}

func toTaskModel(task *dto.TaskDTO) (*models.Task, error) {
	taskId, err := models.BuildTaskId(task.Id)
	if err != nil {
		return nil, err
	}
	return models.BuildTask(taskId, task.Name)
}
