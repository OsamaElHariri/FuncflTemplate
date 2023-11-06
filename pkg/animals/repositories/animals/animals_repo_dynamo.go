package animals

import (
	"funcfltemplate/pkg/animals/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type AnimalsRepoDynamo struct {
	TableName string
	Client    *dynamodb.DynamoDB
}

func (repo *AnimalsRepoDynamo) PutAnimal(entity models.Animal) error {
	entityDb := fromAppEntity(entity)
	dbPayload, err := dynamodbattribute.MarshalMap(entityDb)
	if err != nil {
		return err
	}

	_, err = repo.Client.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(repo.TableName),
		Item:      dbPayload,
	})
	return err
}

func (repo *AnimalsRepoDynamo) GetAnimal(id string) (models.Animal, bool, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(id),
			},
			"SK": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(repo.TableName),
	}

	result, err := repo.Client.GetItem(input)
	if err != nil {
		return models.Animal{}, false, err
	}

	if len(result.Item) == 0 {
		return models.Animal{}, false, nil
	}

	var entityDb AnimalDb
	err = dynamodbattribute.UnmarshalMap(result.Item, &entityDb)
	if err != nil {
		return models.Animal{}, false, err
	}

	return toAppEntity(entityDb), true, nil
}
