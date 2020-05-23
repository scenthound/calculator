package main

import (
	"fmt"
	"bytes"
	"context"
	b64 "encoding/base64"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	pb "scenthound/calculator/calculator"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequiest) (events.APIGatewayProxyResponse, error) {
	// Decode the base64 encoded message
}

func main() {
	lambda.Start(Handler)
}

