package main

import (
	"context"
	"log"

	pb "github.com/hiiamtrong/go-grpc-example/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {

	r, err := c.SayHello(context.Background(), &pb.GreetRequest{Name: "Trong"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(r.Result)
}
