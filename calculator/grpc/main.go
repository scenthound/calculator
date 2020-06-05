package main

import (
	"context"
	b64 "encoding/base64"

	"github.com/golang/protobuf/proto"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	calculator "scenthound/calculator/calculator"
	pb "scenthound/calculator/proto"

)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody []byte

	// Create a calculator instance from the protobuf interface
	var cs pb.CalculatorServer = calculator.Calculator{}

	// Decode the base64 encoded message
	if req.IsBase64Encoded {
		requestBody, _ = b64.StdEncoding.DecodeString(req.Body)
	}

	// Unmarshal back to protobuf
	expression := &pb.Expression{}
	if err := proto.Unmarshal(requestBody, expression); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404}, err
	}

	// Create protobuf response object
	response, err := cs.Evaluate(ctx, expression)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404}, err
	}

	// Marshal response to binary
	responseBody, err := proto.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404}, err
	}

	// Ensure the body is Base64 encoded and Content-Type is set to 'application/x-protobuf'
	resp := events.APIGatewayProxyResponse{
		StatusCode: 200,
		IsBase64Encoded: true,
		Body: b64.StdEncoding.EncodeToString(responseBody),
		Headers: map[string]string {
			 "Content-Type": "application/x-protobuf",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
