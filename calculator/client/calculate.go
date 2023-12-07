package main

import (
	"context"
	"log"

	pb "github.com/l1s4un/go-gRPC/calculator/proto"
)

func doCalculate(c pb.CalculatorServiceClient) {
	log.Println("doCalculate was invoked")

	res, err := c.Calculate(context.Background(), &pb.CalculatorRequest{
		A: 5,
		B: 10,
	})
	if err != nil {
		log.Fatalf("Could not calculate $v\n", err)
	}

	log.Printf("Calculating result: %v\n", res.Sum)
}
