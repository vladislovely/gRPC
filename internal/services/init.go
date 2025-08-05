package services

import (
	"google.golang.org/grpc"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"

	pb "vladislove-gRPC/internal/gen/user"
	repo "vladislove-gRPC/internal/infrastructure/database/repository"
	UserService "vladislove-gRPC/internal/services/user"
)

func RegisterServices(server *grpc.Server, logger *log.Logger, db *gorm.DB) {
	userRepo := repo.NewUserRepository(db)

	pb.RegisterUserServiceServer(server, UserService.NewUserServiceServer(logger, userRepo))
}
