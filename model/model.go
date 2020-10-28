package model

import (
	"time"
)

/*
Base struct is used to give some properties to all other 
structs that import it (as a template). Certainly, 
User and Task structs take the Base to import its 
properties
*/
type Base struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Task struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      uint   `json:"userid"`
}

type User struct {
	Base
	Username string `json:"username"`
	Password string `json:"password"`
	Todos    []Task `json:"todos"`
}
