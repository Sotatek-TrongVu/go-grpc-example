package main

import (
	"context"
	"log"

	pb "github.com/hiiamtrong/go-grpc-example/calculator/proto"
)

func Sum(c pb.CalculatorServiceClient) {
	rs, err := c.Sum(context.Background(), &pb.SumRequest{FirstNumber: 2, SecondNumber: 2})

	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}

	log.Println(rs.Result)
}
