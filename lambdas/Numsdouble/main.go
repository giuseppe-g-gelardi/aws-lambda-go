package main

import (
	"encoding/json"
	"net/http"

	"aws-lambda-go/src/util"

	"github.com/aws/aws-lambda-go/lambda"
)

type RequestBody struct {
	Num int `json:"num"`
}

type ResponseBody struct {
	Message string `json:"message"`
	Value   int    `json:"value"`
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

	double := requestBody.Num * 2

	responseBody := ResponseBody{
		Message: "The value of the number you sent * 2 is: ",
		Value:   double,
	}

	responseJSON, err := json.Marshal(responseBody)
	if err != nil {
		return util.Response{
			StatusCode: http.StatusInternalServerError,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       `{"error": "internal server error}`,
		}, nil
	}

	// res := lambda_response(responseJSON)
	response := util.Response{
		Body:       string(responseJSON),
		StatusCode: http.StatusOK,
		Headers:    util.DefaultHeaders(),
	}
	return response, nil
	// return res, nil
}

// func lambda_response(json []byte) util.Response {
// 	return util.Response{
// 		Body:       string(json),
// 		StatusCode: http.StatusOK,
// 		Headers:    util.DefaultHeaders(),
// 	}
// }
