package user

import (
	"context"

	log "github.com/sirupsen/logrus"

	pb "vladislove-gRPC/internal/gen/user"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer

	logger *log.Logger
}

func (s *userServiceServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	s.logger.Info("Получен запрос на пользователя с ID: %s", req.Id)

	return &pb.UserResponse{
		Id:    req.Id,
		Name:  "Иван Иванов",
		Email: "ivan@example.com",
	}, nil
}

func NewUserServiceServer(logger *log.Logger) pb.UserServiceServer {
	return &userServiceServer{
		logger: logger,
	}
}
