package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "vladislove-gRPC/internal/gen/user"
)

func main() {
	conn, err := grpc.NewClient("server:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetUser(ctx, &pb.UserRequest{Id: "123"})
	if err != nil {
		log.Fatalf("Ошибка при вызове GetUser: %v", err)
	}

	log.Printf("Получены данные о пользователе: ID=%s, Name=%s, Email=%s", res.Id, res.Name, res.Email)
}
