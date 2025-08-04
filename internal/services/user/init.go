package user

import (
	"context"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	pb "vladislove-gRPC/internal/gen/user"
	repo "vladislove-gRPC/internal/infrastructure/database/repository"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer

	logger     *log.Logger
	repository repo.UserRepository
}

func (s *userServiceServer) GetUser(
	ctx context.Context,
	req *pb.UserRequest,
) (*pb.UserResponse, error) {
	s.logger.Printf("Получен запрос на пользователя с ID: %s", req.GetId())

	id, parseUUID := uuid.Parse(req.GetId())
	if parseUUID != nil {
		s.logger.Error("Ошибка парсинга UUID: ", parseUUID)
		return nil, parseUUID
	}

	user, err := s.repository.Get(ctx, id)
	if err != nil {
		s.logger.Error("Ошибка получения пользователя: ", err)
		return nil, err
	}

	return &pb.UserResponse{
		Id:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func NewUserServiceServer(logger *log.Logger, repo repo.UserRepository) pb.UserServiceServer {
	return &userServiceServer{
		logger:     logger,
		repository: repo,
	}
}
