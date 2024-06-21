package env

import (
	"fmt"
	"log"
	"message-service/pkg/utils"
	"os"
)

var (
	GIN_MODE      = ""
	VAULT_ADDRESS = ""
	VAULT_TOKEN   = ""
	SERVICE_NAME  = "External service"
	VERSION       = ""
)

func InitalizeEnvs() {
	env := os.Getenv("ENV")
	vaultAddress := os.Getenv("VAULT_ADDRESS")
	vaultToken := os.Getenv("VAULT_TOKEN")

	if vaultAddress == "" {
		log.Fatal("GetCurrentEnv", "current vault address is not set, please set it. run 'export VAULT_ADDRESS=<VAULT_ADDRESS>'", nil)
	}

	if vaultToken == "" {
		log.Fatal("GetCurrentEnv", "current vault token is not set, please set it. run 'export VAULT_TOKEN=<VAULT_TOKEN>'", nil)
	}

	if env == "" {
		utils.LogWarn("GetCurrentEnv", "current env is not set, defaulting to dev", nil)
		env = "dev"
	}

	version, err := os.ReadFile("./.version")
	if err != nil {
		utils.LogError("GetCurrentEnv", "error reading .version file. make sure it's read, otherwise version will be set to 'unknown", err)
		VERSION = "unknown"
	} else {
		VERSION = fmt.Sprint(string(version))
	}

	switch env {
	case "prod":
		GIN_MODE = "release"
	default:
		GIN_MODE = "debug"
	}

	VAULT_ADDRESS = vaultAddress
	VAULT_TOKEN = vaultToken
}
