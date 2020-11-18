package main

import (
	"net"
	"google.golang.org/grpc"
	"github.com/hashicorp/go-hclog"
	"github.com/jkarlos000/sc-market/provider/internal/infrastructure/server"
	"github.com/jkarlos000/sc-market/provider/internal/core/services/providerssrv"
	"github.com/jkarlos000/sc-market/provider/internal/infrastructure/repositories/gorm"
	"github.com/jkarlos000/sc-market/provider/internal/infrastructure/delivery/grpc/proto"
)

type Logger struct {
	l hclog.Logger
}

func main() {
	//opts := grpc.WithInsecure()
	logger := hclog.Default()
	logger.Info("Iniciando servicio en puerto 5021")
	lis, err := net.Listen("tcp", ":5021")
	if err != nil {
		logger.Error("Failed to listen","error", err)
	}
	logger.Info("Configurando servicios...")
	repository := gorm.NewProvidersRepository("localhost", "providersvc", "OQRqxaRzsd1AWVZ6", "erp_provider", "America/La_Paz", 5432)
	service := providerssrv.NewService(repository)
	srv := server.NewServer(service, logger)
	gserver := grpc.NewServer()
	logger.Info("Levantando servicio GRPC.")
	proto.RegisterProviderServiceServer(gserver, srv)
	if err := gserver.Serve(lis); err != nil {
		logger.Error("Failed to server","error", err)
	}
}