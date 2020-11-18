package main

import (
	"github.com/joho/godotenv"
	"net"
	"os"
	"google.golang.org/grpc"
	"github.com/hashicorp/go-hclog"
	"github.com/jkarlos000/sc-market/provider/internal/infrastructure/server"
	"github.com/jkarlos000/sc-market/provider/internal/core/services/providerssrv"
	"github.com/jkarlos000/sc-market/provider/internal/infrastructure/repositories/gorm"
	"github.com/jkarlos000/sc-market/provider/internal/infrastructure/delivery/grpc/proto"
	"strconv"
)

type Logger struct {
	l hclog.Logger
}

func main() {
	//opts := grpc.WithInsecure()
	logger := hclog.Default()
	logger.Info("Iniciando servicio en puerto "+Config("APP_PORT"))
	lis, err := net.Listen("tcp", ":"+Config("APP_PORT"))
	if err != nil {
		logger.Error("Failed to listen","error", err)
	}
	logger.Info("Configurando servicios...")
	portdb, _ := strconv.ParseUint(Config("DB_PORT"), 10, 32)
	repository := gorm.NewProvidersRepository(Config("DB_HOST"), Config("DB_USER"), Config("DB_PASSWORD"), Config("DB_NAME"), Config("DB_TIMEZONE"), int(portdb))
	service := providerssrv.NewService(repository)
	srv := server.NewServer(service, logger)
	gserver := grpc.NewServer()
	logger.Info("Levantando servicio GRPC.")
	proto.RegisterProviderServiceServer(gserver, srv)
	if err := gserver.Serve(lis); err != nil {
		logger.Error("Failed to server","error", err)
	}
}

func Config(key string) string {
	l := hclog.Default()
	err := godotenv.Load(".env")
	if err != nil {
		l.Error("Config Falla al cargar configuración", "error", err)
	}
	return os.Getenv(key)
}
