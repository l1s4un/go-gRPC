package main

import (
	"context"
	"log"

	pb "github.com/l1s4un/go-gRPC/calculator/proto"
)

func (s *Server) Calculate(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculate function was invoked with %v\n", in)
	return &pb.CalculatorResponse{
		Sum: in.A + in.B,
	}, nil
}
