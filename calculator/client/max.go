package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/hiiamtrong/go-grpc-example/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	arr := []int32{
		2, 5, 1, 10, 2, 30, 25, 49,
	}

	waitc := make(chan struct{})
	go func() {
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving response: %v", err)
			}
			log.Printf("Response from server: %v", req.GetMax())
		}

		close(waitc)
	}()

	go func() {
		for _, i := range arr {
			stream.Send(&pb.MaxRequest{Number: i})
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	<-waitc
}
