package vault

import (
	"fmt"
	"log"
	"os"

	vaultclient "github.com/hashicorp/vault/api"
)

var (
	Envars map[string]interface{}
)

func init() {
	client, err := vaultclient.NewClient(&vaultclient.Config{
		Address: fmt.Sprintf("%s:8200", os.Getenv("VAULT_ADDRESS")),
	})
	if err != nil {
		log.Fatalf("Error initializing Vault client: %s", err)
	}

	vaultToken := os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		log.Fatal("VAULT_TOKEN is not set")
	}
	client.SetToken(vaultToken)

	path := "pickside/data/credentials"

	secret, err := client.Logical().Read(path)
	if err != nil {
		log.Fatalf("Error reading secret: %s", err)
	}

	if secret == nil || secret.Data == nil {
		log.Fatal("Secret data is nil. Ensure the secret path is correct and the secret exists.")
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		log.Fatal("Secret structure is not as expected. Unable to find 'data' map.")
	}

	log.Println("Service is connected to Vault.")

	Envars = data
}
