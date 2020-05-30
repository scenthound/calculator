package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/golang/protobuf/proto"

	"scenthound/cfwrap"
	pb "scenthound/calculator/calculator-proto"
)

func main() {

	// Get the API GW calculator URL from cloudformation
	calcUrl, err := cfwrap.GetExport("us-east-1", "dev-calculator-sayHello")

	if err != nil {
		log.Println(err)
	}

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
	request, err := http.NewRequest("POST", calcUrl, bytes.NewBuffer(requestBody))
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
