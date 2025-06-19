package app

import (
	"GonIO/internal/domain"
	envzilla "GonIO/pkg/EnvZilla"
	csvparser "GonIO/pkg/myCSV"
	"encoding/csv"
	"errors"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
)

func init() {
	slog.Info("Starting config loading...")
	if err := envzilla.Loader("configs/.env"); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			slog.Warn("Config file is not exist...")
			slog.Info("Start reading cmd arguments")

		} else {
			log.Fatal("Configs loading error: ", err)
		}
	}

	if err := ParseConfig(); err != nil {
		log.Fatal("Config validation error: ", err)
	}
	slog.Info("Config loading finished...")

	slog.Info("Metadata file check...")

	CheckDir()
	CreateMetaData()

	slog.Info("Everything is OK...")
}

func ParseConfig() error {
	domain.Port = os.Getenv("PORT")
	domain.Host = os.Getenv("HOST")
	domain.BucketsPath = os.Getenv("BUCKETPATH")

	if len(domain.Host) == 0 {
		return domain.ErrEmptyDomain
	}

	portInt, err := strconv.Atoi(domain.Port)
	if err != nil {
		slog.Debug("Port convert error: ", "portNum", portInt, "error", "invalid port number")
		return domain.ErrInvalidPortStr
	}

	if portInt < 1100 || portInt > 65535 {
		return domain.ErrInvalidPortStr
	}

	if domain.BucketsPath == "" {
		domain.BucketsPath = "data"
	}

	return nil
}

func CreateMetaData() {
	data := []string{"Name", "CreationTime", "LastModifiedTime", "Status"}
	domain.BucketsMetaPath = domain.BucketsPath + "/buckets.csv"

	empty, err := csvparser.CheckEmpty(domain.BucketsMetaPath)
	if err != nil {
		log.Fatal("Failed to read bucket metadata : ", err.Error())
	}

	if !empty {
		return
	}

	file, err := os.OpenFile(domain.BucketsMetaPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		log.Fatal("Failed to create bucket metadata: ", err.Error())
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(data)
	if err != nil {
		log.Fatal("Failed to write CSV metadata: ", err.Error())
	}
}

func CheckDir() {
	absPath, err := filepath.Abs(domain.BucketsPath)
	if err != nil {
		log.Fatal("Error resolving absolute path:", err)
	}

	_, err = os.Stat(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(absPath, os.ModePerm)
			if err != nil {
				log.Fatal("Error create directory :", err)
			}
		} else {
			log.Fatal("Error checking path:", err)
		}
	}
}
