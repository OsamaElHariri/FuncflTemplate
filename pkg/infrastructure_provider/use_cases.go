package infrastructure_provider

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	instance *session.Session
)

func getSession() *session.Session {
	if instance == nil {
		instance = session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
	}
	return instance
}

func GetTableName() string {
	// replace this by the name of the table provided for you
	return "2X1AfLMD6qumGJ3pjliWIwfyEF9-MyFirtProject"
}

func GetDynamo() *dynamodb.DynamoDB {
	return dynamodb.New(getSession())
}
