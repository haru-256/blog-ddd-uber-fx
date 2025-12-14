package dto

type TaskDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserDTO struct {
	Id    string     `json:"id"`
	Name  string     `json:"name"`
	Tasks []*TaskDTO `json:"tasks"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

type GetUserByIdRequest struct {
	Id string
}

type UpdateUserRequest struct {
	Id    string
	Name  string     `json:"name"`
	Tasks []*TaskDTO `json:"tasks"`
}

type DeleteUserRequest struct {
	Id string
}

type CreateTaskRequest struct {
	Name string `json:"name"`
}

type GetTaskByIdRequest struct {
	Id string
}

type UpdateTaskRequest struct {
	Id   string
	Name string `json:"name"`
}

type DeleteTaskRequest struct {
	Id string
}
