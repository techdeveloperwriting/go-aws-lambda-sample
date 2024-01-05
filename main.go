package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(_ context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Hello Handler")
	data, _ := json.Marshal(map[string]interface{}{
		"message": "hello world Golang new",
		"event":   r,
	})
	log.Println("Hello Handler Data", string(data))
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(data),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
