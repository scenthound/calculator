package calculator

import (
	context "context"
	pb "scenthound/calculator/proto"
)

type Calculator struct {}

// Implement the interface exposed by the compiled protobuf
// For other languages, this would typically involve extending the serivce, and
// overriding the rpc method
func (c Calculator) Evaluate(ctx context.Context, req *pb.Expression) (*pb.Response, error) {

	var result float64

	switch req.Operation {
		case pb.Operations_ADD:
			result = req.Operands.Number1 + req.Operands.Number2
		case pb.Operations_SUBTRACT:
			result = req.Operands.Number1 - req.Operands.Number2
		case pb.Operations_MULTIPLY:
			result = req.Operands.Number1 * req.Operands.Number2
		case pb.Operations_DIVIDE:
			result = req.Operands.Number1 / req.Operands.Number2
	}

	return &pb.Response {
			Expression: req,
			Result: result,
		}, nil
}
