package main

import (
	"github.com/Suhach/task-service/internal/client"
	"github.com/Suhach/task-service/internal/database"
	"github.com/Suhach/task-service/internal/task"
	"github.com/Suhach/task-service/internal/transport/grpc"
	"log"
)

func main() {
	database.InitDB()

	userClient, err := client.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}

	repo := task.NewRepository(database.DB)
	svc := task.NewService(repo, userClient)
	handler := grpc.NewHandler(svc)

	grpc.RunServer(handler, ":50052")
}
