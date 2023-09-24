package config

import "os"

type TodoConfig struct {
	Address string
}

type UserConfig struct {
	Address string
}

func NewTodoConfig() *TodoConfig {
	cfg := &TodoConfig{
		Address: LookUpEnv("TODO_SERVER_ADDRESS", "todo_server:8000"),
	}
	return cfg
}

func NewUserConfig() *UserConfig {
	cfg := &UserConfig{
		Address: LookUpEnv("USER_SERVER_ADDRESS", "user_server:8000"),
	}
	return cfg
}

func LookUpEnv(key string, fallback string) string {
	if e, ok := os.LookupEnv(key); ok {
		return e
	} else {
		return fallback
	}
}
