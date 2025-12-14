package models

import (
	"errors"
)

type Task struct {
	id   *TaskId
	name string
}

type User struct {
	id    *UserId
	name  string
	tasks []*Task
}

func NewUser(name string, tasks []*Task) (*User, error) {
	user := &User{
		id:    NewUserId(),
		tasks: tasks,
	}
	if err := user.SetName(name); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Id() *UserId {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetId(id *UserId) error {
	u.id = id
	return nil
}

func (u *User) SetName(name string) error {
	if len(name) == 0 {
		return errors.New("name is empty")
	}
	u.name = name
	return nil
}

func (u *User) Tasks() []*Task {
	return u.tasks
}

func (u *User) SetTasks(tasks []*Task) {
	u.tasks = tasks
}

func BuildUser(id *UserId, name string, tasks []*Task) (*User, error) {
	user := new(User)
	user.SetTasks(tasks)
	if err := user.SetId(id); err != nil {
		return nil, err
	}
	if err := user.SetName(name); err != nil {
		return nil, err
	}
	return user, nil
}

func NewTask(name string) (*Task, error) {
	if len(name) == 0 {
		return nil, errors.New("name is empty")
	}
	return &Task{
		id:   NewTaskId(),
		name: name,
	}, nil
}

func (t *Task) Id() *TaskId {
	return t.id
}

func (t *Task) SetId(id *TaskId) error {
	t.id = id
	return nil
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) SetName(name string) error {
	if len(name) == 0 {
		return errors.New("name is empty")
	}
	t.name = name
	return nil
}

func BuildTask(id *TaskId, name string) (*Task, error) {
	task := new(Task)
	if err := task.SetId(id); err != nil {
		return nil, err
	}
	if err := task.SetName(name); err != nil {
		return nil, err
	}
	return task, nil
}
