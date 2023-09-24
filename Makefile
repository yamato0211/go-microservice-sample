user_evans: ## Run evans.
	evans --proto ./user/proto/user.proto --port 8001

todo_evans:
	evans --proto ./todo/proto/todo.proto --port 8002

gqlgen:
	go run github.com/99designs/gqlgen generate

dry-fix:
	golangci-lint run ./...

fix:
	golangci-lint run --fix