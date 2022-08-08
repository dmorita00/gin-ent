package dotenvutil

import (
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/xerrors"
)

func Init() error {
	file, err := getEnvFile()
	if err != nil {
		return xerrors.Errorf("message: %w", err)
	}
	err = godotenv.Load(file)
	if err != nil {
		return xerrors.Errorf("message: %w", err)
	}
	return nil
}

func getEnvFile() (string, error) {
	dir, _ := os.Getwd()
	var file string
	var err error
	file = dir + "/.env"
	_, err = os.Stat(file)
	if err != nil {
		return "", err
	}
	return file, nil
}
