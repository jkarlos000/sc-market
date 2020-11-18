package main

import (
	"github.com/jkarlos000/sc-market/user/internal/core/services/userssrv"
	"github.com/jkarlos000/sc-market/user/internal/infrastructure/delivery/grpc/proto"
	"github.com/jkarlos000/sc-market/user/internal/infrastructure/repositories/gorm"
	"github.com/jkarlos000/sc-market/user/internal/infrastructure/server"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"github.com/hashicorp/go-hclog"
	"log"
	"net"
	"strconv"
	"os"
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
		log.Fatalf("Failed to listen on %v", err)
	}
	logger.Info("Configurando servicios...")
	portdb, _ := strconv.ParseUint(Config("DB_PORT"), 10, 32)
	repository := gorm.NewUsersRepository(Config("DB_HOST"), Config("DB_USER"), Config("DB_PASSWORD"), Config("DB_NAME"), Config("DB_TIMEZONE"), int(portdb))
	service := userssrv.NewService(repository)
	srv := server.NewServer(service, logger)
	gserver := grpc.NewServer()
	logger.Info("Levantando servicio GRPC.")
	proto.RegisterUserServiceServer(gserver, srv)
	if err := gserver.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
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
