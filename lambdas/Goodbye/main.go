package main

import (
	"encoding/json"
	"net/http"

	"aws-lambda-go/src/util"

	"github.com/aws/aws-lambda-go/lambda"
)

type RequestBody struct {
	Name string `json:"name"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(request util.Request) (util.Response, error) {
	var requestBody RequestBody

	if err := json.Unmarshal([]byte(request.Body), &requestBody); err != nil {
		return util.Response{
			StatusCode: http.StatusBadRequest,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       `{"error": "invalid request body"}`,
		}, nil
	}

	responseBody := map[string]string{
		"message": "Goodbye, " + requestBody.Name + "!",
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
