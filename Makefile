run:
	go run cmd/auth/main.go &
	go run cmd/user/main.go &
	go run cmd/chat/main.go &
	go run cmd/api-gateway/main.go &
	wait
