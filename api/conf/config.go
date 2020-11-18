package conf

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
)

func Config(key string) string {
	l := hclog.Default()
	err := godotenv.Load("./env")
	if err != nil {
		l.Error("Config Falla al cargar configuración", "error", err)
	}
	return os.Getenv(key)
}
