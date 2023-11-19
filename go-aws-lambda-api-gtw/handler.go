package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}

const (
	GET_RAW_PATH    = "/getPerson"
	CREATE_RAW_PATH = "/createPerson"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(event)
	if event.Path == GET_RAW_PATH {
		fmt.Println("Start request for GetPerson")
		personId := event.QueryStringParameters["personId"]
		fmt.Println("Received request for", personId)
	} else if event.Path == CREATE_RAW_PATH {
		fmt.Println("Start request for CreatePerson")
	}
	var person Person
	err := json.Unmarshal([]byte(event.Body), &person)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error na conversão",
		}, nil
	}
	personResponse, err := json.Marshal(person)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error na conversão",
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(personResponse),
	}, nil
}
