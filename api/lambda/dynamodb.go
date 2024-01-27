package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetCounter(svc *dynamodb.DynamoDB, tableName, id string) (Item, error) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return Item{}, err
	}
	if result.Item == nil {
		return Item{}, fmt.Errorf("Could not find item with id = 1")
	}
	var item Item
	println(result.String())
	println(result.Item)
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return Item{}, err
	}
	fmt.Println("Found visitor")
	fmt.Println("Visitor number: ", item.Views)
	return item, nil
}

func updateCounter(Item Item, svc *dynamodb.DynamoDB, tableName string) (Item, error) {
	Item.Views = Item.Views + 1
	av, err := dynamodbattribute.MarshalMap(Item)
	if err != nil {
		return Item, err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		return Item, err
	}
	return Item, nil
}
