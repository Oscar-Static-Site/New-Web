package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Item struct {
	Views int    `json:"views"`
	Id    string `json:"id"`
}

func counterops() Item {
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
	return item
}

func handler(
	request events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	item := counterops()
	responsebody := map[string]string{
		"message": string(item.Views),
	}
	responeJson, err := json.Marshal(responsebody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: `{"error": "Could not marshal response"}`,
		}, nil
	}
	respone := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responeJson),
		Headers: map[string]string{
			"Content-Type":                     "text/plain",
			"Access-Control-Allow-Origin":      "*.oscarcorner.com",
			"Access-Control-Allow-Headers":     "Content-Type",
			"Access-Control-Allow-Methods":     "GET,OPTIONS,POST,PUT,DELETE",
			"Access-Control-Allow-Credentials": "true",
		},
	}
	return respone, nil
}

func main() {
	lambda.Start(handler)
}
