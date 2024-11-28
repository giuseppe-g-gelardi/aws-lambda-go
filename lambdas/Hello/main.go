package main

import (
	"encoding/json"
	"net/http"

	"aws-lambda-go/src/util"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

func Handler(request util.Request) (util.Response, error) {
	responseBody := map[string]string{
		"message": "some kind of message",
	}

	responseJSON, err := json.Marshal(responseBody)
	if err != nil {
		return util.Response{
			StatusCode: http.StatusInternalServerError,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       `{"error": "internal server error}`,
		}, nil
	}

	response := util.Response{
		Body:       string(responseJSON),
		StatusCode: http.StatusOK,
		Headers:    util.DefaultHeaders(),
	}
	return response, nil
}
