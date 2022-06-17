package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/hiiamtrong/go-grpc-example/greet/proto"
)

func doGreetLong(c pb.GreetServiceClient) {
	reqs := []*pb.GreetRequest{
		{
			Name: "Trong",
		},
		{
			Name: "Dep"},
		{
			Name: "Trai"},
	}

	stream, err := c.SayLongTime(context.Background())

	if err != nil {
		log.Fatalf("Error while calling SayLongTime RPC: %v", err)
	}

	for _, req := range reqs {
		fmt.Printf("Sending req: %v\n", req)

		stream.Send(req)

		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)

	}

	log.Printf("Response from server: %v", res.GetResult())
}
