package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "vladislove-gRPC/internal/gen/user"
	boot "vladislove-gRPC/internal/infrastructure/bootstrap"
)

func main() {
	cfg, bootstrapErr := boot.ConfigFromEnv()
	if bootstrapErr != nil {
		log.Fatalf("произошла ошибка при прочтении конфигурации .env: %v", bootstrapErr)
	}

	target := fmt.Sprintf("%s:%d", cfg.GRPCAddr, cfg.GRPCPort)

	conn, clientErr := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if clientErr != nil {
		log.Fatalf("Ошибка подключения: %v", clientErr)
	}
	defer func() { _ = conn.Close() }()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, getUserErr := client.GetUser(
		ctx,
		&pb.UserRequest{Id: "6efa6f6e-60ac-466e-8561-1a5f56e1f2cb"},
	)

	if getUserErr != nil {
		cancel()

		log.Fatalf("Ошибка при вызове GetUser: %v", getUserErr)
	}

	log.Printf(
		"Получены данные о пользователе: ID=%s, Name=%s, Email=%s",
		res.GetId(),
		res.GetName(),
		res.GetEmail(),
	)
}
