cli:
	go build -mod vendor -ldflags="-s -w" -o bin/ecs-launch-task cmd/ecs-launch-task/main.go
