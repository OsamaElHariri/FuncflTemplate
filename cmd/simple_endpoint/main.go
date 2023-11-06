package main

import (
	"context"
	"encoding/json"
	"funcfltemplate/pkg/animals"
	networkcontracts "funcfltemplate/pkg/network_contracts"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event *events.APIGatewayProxyRequest) (events.APIGatewayV2HTTPResponse, error) {
	usecases := animals.NewUseCases()

	result := usecases.MakeSound()

	stringBody, _ := json.Marshal(networkcontracts.MakeSoundResponse{
		Sound: result,
	})

	return events.APIGatewayV2HTTPResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: 200,
		Body:       string(stringBody),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
