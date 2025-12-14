package server

import (
	"log/slog"
	"net/http"

	"github.com/haru-256/blog-ddd-uber-fx/internal/application/dto"
	"github.com/haru-256/blog-ddd-uber-fx/internal/application/service"
	"github.com/labstack/echo/v4"
)

// RouteRegistrar はEchoへのルート登録を行うインターフェースです。
// Group機能を使って一括登録するために使用します。
type RouteRegistrar interface {
	Register(e *echo.Echo)
}

type UserServiceHandler struct {
	logger *slog.Logger
	svc    service.UserService
}

func NewUserServiceHandler(logger *slog.Logger, svc service.UserService) *UserServiceHandler {
	return &UserServiceHandler{
		logger: logger,
		svc:    svc,
	}
}

// Register はUserServiceのルートを登録します。
func (h *UserServiceHandler) Register(e *echo.Echo) {
	g := e.Group("/users")
	g.POST("", h.CreateUser)
	g.GET("", h.GetAllUsers)
	g.GET("/:id", h.GetUserById)
	g.PUT("/:id", h.UpdateUser)
	g.DELETE("/:id", h.DeleteUser)
}

func (h *UserServiceHandler) CreateUser(c echo.Context) error {
	req := new(dto.CreateUserRequest)
	if err := c.Bind(req); err != nil {
		h.logger.Error("Failed to bind request", "error", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := h.svc.Create(c.Request().Context(), req.Name)
	if err != nil {
		h.logger.Error("Failed to create user", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserServiceHandler) GetAllUsers(c echo.Context) error {
	users, err := h.svc.GetAll(c.Request().Context())
	if err != nil {
		h.logger.Error("Failed to get all users", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserServiceHandler) GetUserById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is required")
	}
	user, err := h.svc.GetById(c.Request().Context(), id)
	if err != nil {
		h.logger.Error("Failed to get user by id", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserServiceHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is required")
	}
	req := new(dto.UpdateUserRequest)
	if err := c.Bind(req); err != nil {
		h.logger.Error("Failed to bind request", "error", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := h.svc.Update(c.Request().Context(), id, req.Name, req.Tasks)
	if err != nil {
		h.logger.Error("Failed to update user", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserServiceHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is required")
	}
	if err := h.svc.Delete(c.Request().Context(), id); err != nil {
		h.logger.Error("Failed to delete user", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

type TaskServiceHandler struct {
	logger *slog.Logger
	svc    service.TaskService
}

func NewTaskServiceHandler(logger *slog.Logger, svc service.TaskService) *TaskServiceHandler {
	return &TaskServiceHandler{
		logger: logger,
		svc:    svc,
	}
}

// Register はTaskServiceのルートを登録します。
func (h *TaskServiceHandler) Register(e *echo.Echo) {
	g := e.Group("/tasks")
	g.POST("", h.CreateTask)
	g.GET("", h.GetAllTasks)
	g.GET("/:id", h.GetTaskById)
	g.PUT("/:id", h.UpdateTask)
	g.DELETE("/:id", h.DeleteTask)
}

func (h *TaskServiceHandler) CreateTask(c echo.Context) error {
	req := new(dto.CreateTaskRequest)
	if err := c.Bind(req); err != nil {
		h.logger.Error("Failed to bind request", "error", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	task, err := h.svc.Create(c.Request().Context(), req.Name)
	if err != nil {
		h.logger.Error("Failed to create task", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskServiceHandler) GetAllTasks(c echo.Context) error {
	tasks, err := h.svc.GetAll(c.Request().Context())
	if err != nil {
		h.logger.Error("Failed to get all tasks", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskServiceHandler) GetTaskById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is required")
	}
	task, err := h.svc.GetById(c.Request().Context(), id)
	if err != nil {
		h.logger.Error("Failed to get task by id", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskServiceHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is required")
	}
	req := new(dto.UpdateTaskRequest)
	if err := c.Bind(req); err != nil {
		h.logger.Error("Failed to bind request", "error", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	task, err := h.svc.Update(c.Request().Context(), id, req.Name)
	if err != nil {
		h.logger.Error("Failed to update task", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskServiceHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id is required")
	}
	if err := h.svc.Delete(c.Request().Context(), id); err != nil {
		h.logger.Error("Failed to delete task", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
