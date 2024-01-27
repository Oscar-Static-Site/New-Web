package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Item struct {
	Views int    `json:"views"`
	Id    string `json:"id"`
}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	tableName := "visitor-counter"
	item, err := GetCounter(svc, tableName, "1")
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	fmt.Println("Found visitor")
	fmt.Println("Visitor number: ", item.Views)
	updateCounter(item, svc, tableName)
}
