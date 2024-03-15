package database

import (
	"fmt"

	"github.com/projects/gambituser/models"
	"github.com/projects/gambituser/tools"
)

func SignUp(signUp models.SignUp) error {
	fmt.Println("Iniciando regristo")

	err := DatabaseConnection()
	if err != nil {
		return err
	}

	defer Database.Close()

	querySQL := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + signUp.UserEmail + "','" + signUp.UserUUID + "','" + tools.DateMySQL() + "')"
	fmt.Println(querySQL)

	_, err = Database.Exec(querySQL)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp Realizado com Sucesso")
	return nil
}
