package main

import (
	"context"
	"io"
	"log"

	pb "github.com/hiiamtrong/go-grpc-example/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	res := &pb.GreetRequest{
		Name: "Trong",
	}
	stream, err := c.SayManyTimes(context.Background(), res)
	if err != nil {
		log.Fatalf("Error while calling SayManyTimes RPC: %v", err)
	}
	for {

		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		log.Println(res.Result)
	}
}
