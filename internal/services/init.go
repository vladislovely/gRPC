package services

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	pb "vladislove-gRPC/internal/gen/user"
	"vladislove-gRPC/internal/infrastructure/database/repository"
	UserService "vladislove-gRPC/internal/services/user"
)

func RegisterServices(server *grpc.Server, logger *log.Logger, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)

	pb.RegisterUserServiceServer(server, UserService.NewUserServiceServer(logger, userRepo))
}
