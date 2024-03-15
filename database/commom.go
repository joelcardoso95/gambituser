package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/projects/gambituser/models"
	"github.com/projects/gambituser/secretmanager"
)

var SecretModel models.SecretRDSJson
var err error
var Database *sql.DB

func ReadSecret() error {
	SecretModel, err = secretmanager.GetSecret(os.Getenv("SecretName"))
	return err
}

func DatabaseConnection() error {
	Database, err = sql.Open("mysql", StringConnection(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Database.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexão com Mysql realizada com sucesso")
	return nil
}

func StringConnection(key models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = key.Username
	authToken = key.Password
	dbEndpoint = key.Host
	dbName = "gambit"
	databaseString := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowClearTextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(databaseString)
	return databaseString
}
