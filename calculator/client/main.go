package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/golang/protobuf/proto"

	"scenthound/cfwrap"
	pb "scenthound/calculator/proto"
)

func main() {

	var opt int
	var num1, num2 float64

    str := `
---------------------------------------------
 Simple Calculator
---------------------------------------------

Choose operation:

1) Add
2) Subtract
3) Multiply
4) Divide
0) Quit
`

	// Get the API GW calculator URL from cloudformation
	calcUrl, err := cfwrap.GetExport("us-east-1", "dev-calculator-sayHello")

	if err != nil {
		log.Println(err)
	}

	// Show the menu
	for {
        fmt.Printf("%s", str)
        fmt.Scanf("%d", &opt)

        if opt == 0 {
            os.Exit(0)
        } else if opt > 4 {
            fmt.Println("Invalid option. Enter an option between (0-4).")
        } else {
            fmt.Println("Enter first number: ")
            fmt.Scanf("%f", &num1)
            fmt.Println("Enter second number: ")
            fmt.Scanf("%f", &num2)
        }

		// Create the protobuf object
		expression := &pb.Expression {
			Operation: pb.Operations(opt-1),
			Operands: &pb.Operands{
				Number1: num1,
				Number2: num2,
			},
		}

		// Marshal our protobuf object into a byte array
		requestBody, err := proto.Marshal(expression)
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

		// Read the response
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		response := &pb.Response{}

		if err := proto.Unmarshal(respBody, response); err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Result: %.2f", response.Result)
	}
}
