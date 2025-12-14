package models

import "github.com/google/uuid"

type UserId struct{ value string }

func NewUserId() *UserId {
	return &UserId{value: uuid.NewString()}
}

func BuildUserId(value string) (*UserId, error) {
	uuidValue, err := uuid.Parse(value)
	if err != nil {
		return nil, err
	}
	return &UserId{value: uuidValue.String()}, nil
}

func (u *UserId) Value() string {
	return u.value
}

func (u *UserId) Equals(other *UserId) bool {
	return u.value == other.value
}

type TaskId struct{ value string }

func NewTaskId() *TaskId {
	return &TaskId{value: uuid.NewString()}
}

func BuildTaskId(value string) (*TaskId, error) {
	uuidValue, err := uuid.Parse(value)
	if err != nil {
		return nil, err
	}
	return &TaskId{value: uuidValue.String()}, nil
}

func (t *TaskId) Value() string {
	return t.value
}

func (t *TaskId) Equals(other *TaskId) bool {
	return t.value == other.value
}
