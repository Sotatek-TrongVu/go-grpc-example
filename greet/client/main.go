package main

import (
	"context"
	"log"

	pb "github.com/hiiamtrong/go-grpc-example/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	addr := "localhost:50051"
	client, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer client.Close()

	c := pb.NewGreetServiceClient(client)

	r, err := c.SayHello(context.Background(), &pb.GreetRequest{Name: "Trong"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(r.Result)
}
