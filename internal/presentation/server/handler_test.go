package server_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/haru-256/blog-ddd-uber-fx/internal/application/dto"
	mock_service "github.com/haru-256/blog-ddd-uber-fx/internal/mocks/service"
	"github.com/haru-256/blog-ddd-uber-fx/internal/presentation/server"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserServiceHandler(t *testing.T) {
	// Setup generic logger for handler
	logger := slog.New(slog.NewJSONHandler(io.Discard, nil))

	t.Run("CreateUser", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSvc := mock_service.NewMockUserService(ctrl)
		h := server.NewUserServiceHandler(logger, mockSvc)

		e := echo.New()
		reqBody := `{"name":"test-user"}`
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedDTO := &dto.UserDTO{
			Id:   "test-id",
			Name: "test-user",
		}

		mockSvc.EXPECT().Create(gomock.Any(), "test-user").Return(expectedDTO, nil)

		if assert.NoError(t, h.CreateUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var got dto.UserDTO
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &got))
			assert.Equal(t, expectedDTO.Id, got.Id)
		}
	})

	t.Run("GetAllUsers", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSvc := mock_service.NewMockUserService(ctrl)
		h := server.NewUserServiceHandler(logger, mockSvc)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedDTOs := []*dto.UserDTO{
			{Id: "1", Name: "u1"},
		}

		mockSvc.EXPECT().GetAll(gomock.Any()).Return(expectedDTOs, nil)

		if assert.NoError(t, h.GetAllUsers(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var got []*dto.UserDTO
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &got))
			assert.Len(t, got, 1)
		}
	})

	t.Run("GetUserById", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSvc := mock_service.NewMockUserService(ctrl)
		h := server.NewUserServiceHandler(logger, mockSvc)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/users/test-id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("test-id")

		expectedDTO := &dto.UserDTO{Id: "test-id", Name: "test"}

		mockSvc.EXPECT().GetById(gomock.Any(), "test-id").Return(expectedDTO, nil)

		if assert.NoError(t, h.GetUserById(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("UpdateUser", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSvc := mock_service.NewMockUserService(ctrl)
		h := server.NewUserServiceHandler(logger, mockSvc)

		e := echo.New()
		reqBody := `{"name":"updated"}`
		req := httptest.NewRequest(http.MethodPut, "/users/test-id", bytes.NewBufferString(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("test-id")

		expectedDTO := &dto.UserDTO{Id: "test-id", Name: "updated"}

		mockSvc.EXPECT().Update(gomock.Any(), "test-id", "updated", gomock.Nil()).Return(expectedDTO, nil)

		if assert.NoError(t, h.UpdateUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("DeleteUser", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSvc := mock_service.NewMockUserService(ctrl)
		h := server.NewUserServiceHandler(logger, mockSvc)

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/users/test-id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("test-id")

		mockSvc.EXPECT().Delete(gomock.Any(), "test-id").Return(nil)

		if assert.NoError(t, h.DeleteUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func TestTaskServiceHandler(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(io.Discard, nil))

	t.Run("CreateTask", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSvc := mock_service.NewMockTaskService(ctrl)
		h := server.NewTaskServiceHandler(logger, mockSvc)

		e := echo.New()
		reqBody := `{"name":"test-task"}`
		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedDTO := &dto.TaskDTO{Id: "t1", Name: "test-task"}
		mockSvc.EXPECT().Create(gomock.Any(), "test-task").Return(expectedDTO, nil)

		if assert.NoError(t, h.CreateTask(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var got dto.TaskDTO
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &got))
			assert.Equal(t, expectedDTO.Id, got.Id)
		}
	})

	t.Run("GetAllTasks", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSvc := mock_service.NewMockTaskService(ctrl)
		h := server.NewTaskServiceHandler(logger, mockSvc)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedDTOs := []*dto.TaskDTO{
			{Id: "t1", Name: "task1"},
		}
		mockSvc.EXPECT().GetAll(gomock.Any()).Return(expectedDTOs, nil)

		if assert.NoError(t, h.GetAllTasks(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var got []*dto.TaskDTO
			assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &got))
			assert.Len(t, got, 1)
		}
	})

	t.Run("GetTaskById", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSvc := mock_service.NewMockTaskService(ctrl)
		h := server.NewTaskServiceHandler(logger, mockSvc)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/tasks/t1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/tasks/:id")
		c.SetParamNames("id")
		c.SetParamValues("t1")

		expectedDTO := &dto.TaskDTO{Id: "t1", Name: "task1"}
		mockSvc.EXPECT().GetById(gomock.Any(), "t1").Return(expectedDTO, nil)

		if assert.NoError(t, h.GetTaskById(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("UpdateTask", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSvc := mock_service.NewMockTaskService(ctrl)
		h := server.NewTaskServiceHandler(logger, mockSvc)

		e := echo.New()
		reqBody := `{"name":"updated-task"}`
		req := httptest.NewRequest(http.MethodPut, "/tasks/t1", bytes.NewBufferString(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/tasks/:id")
		c.SetParamNames("id")
		c.SetParamValues("t1")

		expectedDTO := &dto.TaskDTO{Id: "t1", Name: "updated-task"}
		mockSvc.EXPECT().Update(gomock.Any(), "t1", "updated-task").Return(expectedDTO, nil)

		if assert.NoError(t, h.UpdateTask(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("DeleteTask", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSvc := mock_service.NewMockTaskService(ctrl)
		h := server.NewTaskServiceHandler(logger, mockSvc)

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/tasks/t1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/tasks/:id")
		c.SetParamNames("id")
		c.SetParamValues("t1")

		mockSvc.EXPECT().Delete(gomock.Any(), "t1").Return(nil)

		if assert.NoError(t, h.DeleteTask(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
