package db

import "github.com/google/uuid"

type Database struct {
	usersDB []*userData
	tasksDB []*taskData
}

type taskData struct {
	id   string
	name string
}

type userData struct {
	id    string
	name  string
	tasks []*taskData
}

func NewDatabase() *Database {
	task1 := &taskData{
		id:   uuid.NewString(),
		name: "task1",
	}
	task2 := &taskData{
		id:   uuid.NewString(),
		name: "task2",
	}
	user1 := &userData{
		id:   uuid.NewString(),
		name: "user1",
		tasks: []*taskData{
			task1,
			task2,
		},
	}
	user2 := &userData{
		id:    uuid.NewString(),
		name:  "user2",
		tasks: []*taskData{},
	}
	tasksDB := []*taskData{
		task1,
		task2,
	}
	usersDB := []*userData{
		user1,
		user2,
	}
	return &Database{
		usersDB: usersDB,
		tasksDB: tasksDB,
	}
}
