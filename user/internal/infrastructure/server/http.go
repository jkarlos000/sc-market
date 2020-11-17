package server

import (
	"github.com/jkarlos000/sc-market/user/internal/core/ports"
	"github.com/jkarlos000/sc-market/user/internal/core/services/userssrv"
	"github.com/jkarlos000/sc-market/user/internal/infrastructure/repositories/gorm"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/jkarlos000/sc-market/user/internal/infrastructure/delivery/grpc/proto"
)

type server struct {
	proto.UnimplementedUserServiceServer
	svc ports.UsersService
}

func RunServer(port string) {
	//opts := grpc.WithInsecure()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on %v", err)
	}
	repository := gorm.NewUsersRepository()
	service := userssrv.NewService(repository)
	srv := &server{
		svc: service,
	}
	gserver := grpc.NewServer()
	proto.RegisterUserServiceServer(gserver, srv)
	if err := gserver.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
