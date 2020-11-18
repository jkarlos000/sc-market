package main

import (
	"net"
	"google.golang.org/grpc"
	"github.com/hashicorp/go-hclog"
	"github.com/jkarlos000/sc-market/client/internal/infrastructure/server"
	"github.com/jkarlos000/sc-market/client/internal/core/services/clientssrv"
	"github.com/jkarlos000/sc-market/client/internal/infrastructure/repositories/gorm"
	"github.com/jkarlos000/sc-market/client/internal/infrastructure/delivery/grpc/proto"
)

type Logger struct {
	l hclog.Logger
}

func main() {
	//opts := grpc.WithInsecure()
	logger := hclog.Default()
	logger.Info("Iniciando servicio en puerto 15011")
	lis, err := net.Listen("tcp", ":15011")
	if err != nil {
		logger.Error("Failed to listen","error", err)
	}
	logger.Info("Configurando servicios...")
	repository := gorm.NewClientsRepository("207.244.255.63", "clientsvc", "UJd2XhJ0XDQWgngc", "erp_client", "America/La_Paz", 5432)
	service := clientssrv.NewService(repository)
	srv := server.NewServer(service, logger)
	gserver := grpc.NewServer()
	logger.Info("Levantando servicio GRPC.")
	proto.RegisterClientServiceServer(gserver, srv)
	if err := gserver.Serve(lis); err != nil {
		logger.Error("Failed to server","error", err)
	}
}