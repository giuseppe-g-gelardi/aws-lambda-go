package util

import "github.com/aws/aws-lambda-go/events"

type (
	Request  = events.APIGatewayProxyRequest
	Response = events.APIGatewayProxyResponse
)

func DefaultHeaders() map[string]string {
	return map[string]string{
		"Content-Type":                     "text/plain",
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Headers":     "Content-Type",
		"Access-Control-Allow-Methods":     "OPTIONS, POST, GET",
		"Access-Control-Allow-Credentials": "true",
	}
}
