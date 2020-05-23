package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"bytes"
//	"encoding/json"
	"github.com/golang/protobuf/proto"
//	b64 "encoding/base64"
	pb "scenthound/calculator/calculator"
)

func main() {

	// Create protobuf object
	helloRequest := &pb.HelloRequest {
		Name: "Hendrix",
	}

	// In theroy, out is now a byte array that we should be able to pass to the post request
	requestBody, err := proto.Marshal(helloRequest)
	if err != nil {
		log.Fatalln(err)
	}

	/*
	// Base64 Encode
	encodedRequest := b64.StdEncoding.EncodeToString(requestBody)
	log.Print(encodedRequest)

	// Base64 Decode
	decodedRequest, _ := b64.StdEncoding.DecodeString(encodedRequest)

	// Unmarshal
	helloagain := &pb.HelloRequest{}
	if err := proto.Unmarshal(decodedRequest, helloagain); err != nil {
		log.Fatalln(err)
	}

	log.Println(helloagain)
	*/

	// Create an http client so we can set custom headers
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


	// Curl the API gateway endpoing
	/*	resp, err := http.Post("https://6xazvglpaj.execute-api.us-east-1.amazonaws.com/dev/grpc", "application/x-protobuf", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("after post")
		log.Fatalln(err)
	}*/

	/* Older JSON Marshaling
	requestBody, err := json.Marshal(map[string]string{"name": "Hendrix",})

	// Curl the API gateway endpoint
	resp, err := http.Post("https://jxzyjaatkk.execute-api.us-east-1.amazonaws.com/dev/hello", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}*/

	defer resp.Body.Close()

	helloReply := &pb.HelloReply{}

	// Grab the message body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("grabbing message body")
		log.Fatalln(err)
	}

	if err := proto.Unmarshal(respBody, helloReply); err != nil {
		log.Println(respBody)
		log.Println("unmarshaling")
		log.Fatalln(err)
	}

	// Print the message body
//	log.Println(string(body))
	log.Println(helloReply)
}
