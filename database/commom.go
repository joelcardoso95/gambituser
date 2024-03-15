package database

import (
	"os"

	"github.com/projects/gambituser/models"
	"github.com/projects/gambituser/secretmanager"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err = secretmanager.GetSecret(os.Getenv("SecretName"))
	return err
}
