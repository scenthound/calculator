package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"bytes"
	"github.com/golang/protobuf/proto"
	pb "scenthound/calculator/calculator"
)

func main() {

	// Create protobuf object
	helloRequest := &pb.HelloRequest {
		Name: "Hendrix",
	}

	// Marshal our protobuf object into a byte array
	requestBody, err := proto.Marshal(helloRequest)
	if err != nil {
		log.Fatalln(err)
	}

	// Create an http client and set custom headers
	client := http.Client{}
	request, err := http.NewRequest("POST", "https://6xazvglpaj.execute-api.us-east-1.amazonaws.com/dev/grpc", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/x-protobuf")
	request.Header.Set("Accept", "application/x-protobuf")

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	helloReply := &pb.HelloReply{}

	// Read the response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if err := proto.Unmarshal(respBody, helloReply); err != nil {
		log.Fatalln(err)
	}

	log.Println(helloReply)
}
