package main

import (
	"fmt"
//	"bytes"
	"context"
	b64 "encoding/base64"
	//"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	pb "scenthound/calculator/calculator"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
//	var buf bytes.Buffer
	//var msg string
	var requestBody []byte

	// Decode the base64 encoded message
	if req.IsBase64Encoded {
		//msg = "Message is Base64Encoded"
		requestBody, _ = b64.StdEncoding.DecodeString(req.Body)
	}// else {
		//msg = "Message is NOT Base64Encoded"
	//}

	// Unmarshal back to protobuf
	helloRequest := &pb.HelloRequest{}
	if err := proto.Unmarshal(requestBody, helloRequest); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404}, err
	}

	// Create protobuf response object
	helloReply := &pb.HelloReply {
		Message: helloRequest.Name + " is a good boy.",
	}

	fmt.Println(helloRequest.Name + " is a good boy.")

	// Marshal response to binary
	responseBody, err := proto.Marshal(helloReply)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404}, err
	}

	fmt.Println(responseBody)

	resp := events.APIGatewayProxyResponse{
		StatusCode: 200,
		IsBase64Encoded: true,
		Body: b64.StdEncoding.EncodeToString(responseBody),
		Headers: map[string]string {
			 "Content-Type": "application/x-protobuf",
		},
	}

/*	body, err := json.Marshal(map[string]interface{}{
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
*/
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}

