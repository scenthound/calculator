package calculator 

import (
	context "context"
	pb "scenthound/calculator/proto"
)

type Greeter struct {}

// Implement the interface exposed by the compiled protobuf
// For other languages, this would typically involve extending the service, and
// overriding the rpc method
func (g Greeter) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply {
			Message: req.Name + " is my special guy.",
		}, nil
}

