package model

// User ...
type User struct {
	FullName string `json:"fullName,omitempty"`
	Email    string `json:"email,omitempty"`
	Lock     bool   `json:"lock,omitempty"`
	Roles    []Role `json:"roles,omitempty"`
}

// Role ...
type Role struct {
	Code        string       `json:"code"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}

// Permission ...
type Permission struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// AuthenticatedUser ...
type AuthenticatedUser struct {
	Email string
	Roles []Role
}

type Task struct {
	Id          int32  `json:"id,omitempty"`
	Title       string `json:"title" validate:"empty=false"`
	Description string `json:"description" validate:"empty=false"`
	DueDate     string `json:"due_date,omitempty"`
	State       string `json:"state" validate:"empty=false"`
}
