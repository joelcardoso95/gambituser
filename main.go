package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/projects/gambituser/awsgo"
	"github.com/projects/gambituser/database"
	"github.com/projects/gambituser/models"
)

func main() {
	lambda.Start(LambdaExecute)
}

func LambdaExecute(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitilizeAWS()

	if !ParametersValidate() {
		fmt.Println("Erro nos parametros, deve enviar 'SecretManager'")
		err := errors.New("Erro nos parametros, deve enviar SecretName")
		return event, err
	}

	var signUpData models.SignUp
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			signUpData.UserEmail = att
			fmt.Println("Email = " + signUpData.UserEmail)
		case "sub":
			signUpData.UserUUID = att
			fmt.Println("UUID = " + signUpData.UserUUID)
		}
	}

	err := database.ReadSecret()
	if err != nil {
		fmt.Println("Erro ao ler o Secret " + err.Error())
		return event, err
	}
}

func ParametersValidate() bool {
	var getParameter bool
	_, getParameter = os.LookupEnv("SecretName")
	return getParameter
}
