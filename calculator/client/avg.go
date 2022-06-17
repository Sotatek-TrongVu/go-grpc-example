package main

import (
	"context"
	"log"
	"time"

	pb "github.com/hiiamtrong/go-grpc-example/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	stream, err := c.Avg(context.Background())
	reqs := []*pb.AvgRequest{
		{
			Number: 1,
		},
		{
			Number: 2,
		},
		{
			Number: 3,
		},
		{
			Number: 4,
		},
	}

	if err != nil {
		log.Fatalf("Error while calling Avg RPC: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}
	log.Printf("Response from server: %v", res.GetAvg())
}
