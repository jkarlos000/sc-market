package main

import (
	"github.com/jkarlos000/sc-market/user/internal/core/services/userssrv"
	"github.com/jkarlos000/sc-market/user/internal/infrastructure/delivery/grpc/proto"
	"github.com/jkarlos000/sc-market/user/internal/infrastructure/repositories/gorm"
	"github.com/jkarlos000/sc-market/user/internal/infrastructure/server"
	"google.golang.org/grpc"
	"github.com/hashicorp/go-hclog"
	"log"
	"net"
)

type Logger struct {
	l hclog.Logger
}

func main() {
	//opts := grpc.WithInsecure()
	logger := hclog.Default()
	logger.Info("Iniciando servicio en puerto 5001")
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen on %v", err)
	}
	logger.Info("Configurando servicios...")
	repository := gorm.NewUsersRepository("localhost", "usersvc", "m7TDiQqO7kb3aEY2", "erp_user", "America/La_Paz", 5432)
	service := userssrv.NewService(repository)
	srv := server.NewServer(service, logger)
	gserver := grpc.NewServer()
	logger.Info("Levantando servicio GRPC.")
	proto.RegisterUserServiceServer(gserver, srv)
	if err := gserver.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
