package services

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "vladislove-gRPC/internal/gen/user"
	UserService "vladislove-gRPC/internal/services/user"
)

func RegisterServices(server *grpc.Server, logger *log.Logger) {
	pb.RegisterUserServiceServer(server, UserService.NewUserServiceServer(logger))
}
