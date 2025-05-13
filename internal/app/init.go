package app

import (
	"GonIO/internal/domain"
	envzilla "GonIO/pkg/EnvZilla"
	"log"
	"log/slog"
	"os"
	"strconv"
)

func init() {
	log.Println("Starting config loading...")
	if err := envzilla.Loader("configs/.env"); err != nil {
		log.Fatalf("Configs loading error: %s", err.Error())
	}

	if err := CheckPort(); err != nil {
		log.Fatalf("Config validation error: %s", err.Error())
	}

	log.Println("Config loading finished...")
}

func CheckPort() error {
	domain.Port = os.Getenv("PORT")
	domain.URLDomain = os.Getenv("DOMAIN")

	if len(domain.URLDomain) == 0 {
		return domain.ErrEmptyDomain
	}

	portInt, err := strconv.Atoi(domain.Port)
	if err != nil {
		slog.Debug("Port convert error: ", "portNum", portInt, "Errmessage", err.Error())
		return domain.ErrInvalidPortStr
	}

	if portInt < 1100 || portInt > 65535 {
		return domain.ErrInvalidPortStr
	}

	return nil
}
