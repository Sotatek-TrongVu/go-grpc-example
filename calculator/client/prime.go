package main

import (
	"context"
	"io"
	"log"

	pb "github.com/hiiamtrong/go-grpc-example/calculator/proto"
)

func Prime(c pb.CalculatorServiceClient) {
	rs, err := c.Prime(context.Background(), &pb.PrimeRequest{Number: 243454470})

	if err != nil {
		log.Fatalf("could not prime: %v", err)

	}

	for {
		res, err := rs.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not prime: %v", err)
		}
		log.Println(res.Prime)
	}
}
