package main

import (
	"net/http"
	"io/ioutil"
	"log"
)

func main() {
	
	// Curl the API gateway endpoint
	resp, err := http.Get("https://jxzyjaatkk.execute-api.us-east-1.amazonaws.com/dev/hello")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// Grab the message body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the message body
	log.Println(string(body))
}
