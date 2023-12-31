// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Title  string `json:"title"`
	UserID int    `json:"userId"`
}

type NewUser struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type UpdateTodo struct {
	ID int `json:"id"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
