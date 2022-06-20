package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/hiiamtrong/go-grpc-example/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	reqs := []*pb.GreetRequest{
		{
			Name: "Trong",
		},
		{
			Name: "Dep"},
		{
			Name: "Trai"},
	}

	waitc := make(chan struct{})

	stream, err := c.SayEveryOne(context.Background())

	if err != nil {
		log.Fatalf("Error while calling SayEveryOne RPC: %v", err)
	}

	go func() {
		for {
			req, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving response: %v", err)
			}

			log.Printf("Response from server: %v", req.GetResult())

		}
		close(waitc)
	}()

	go func() {
		for _, req := range reqs {
			log.Printf("Sending req: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)

		}
		stream.CloseSend()
	}()

	<-waitc
}
