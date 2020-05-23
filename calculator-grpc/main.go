package main

import (
//	"fmt"
	"bytes"
	"context"
//	b64 "encoding/base64"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
//	pb "scenthound/calculator/calculator"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var buf bytes.Buffer
	var msg string

	if req.IsBase64Encoded {
		msg = "Message is Base64Encoded"
	} else {
		msg = "Message is NOT Base64Encoded"
	}

	body, err := json.Marshal(map[string]interface{}{
		"message": msg,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := events.APIGatewayProxyResponse{
		StatusCode:	 200,
		IsBase64Encoded: false,
		Body:		 buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
	// Decode the base64 encoded message
}

func main() {
	lambda.Start(Handler)
}

