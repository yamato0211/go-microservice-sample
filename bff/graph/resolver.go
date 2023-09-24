package graph

import "go-micro-sample/bff/client"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService client.IUserService
	TodoService client.ITodoService
}
