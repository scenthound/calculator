package main

import (
	"fmt"
	pb "scenthound/calculator/calculator"
)

func main() {
	p := pb.HelloRequest{
		Name: "Hendrix",
	}

	_ = p
	fmt.Println("hello")
}
