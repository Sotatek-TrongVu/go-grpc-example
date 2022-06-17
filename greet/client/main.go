package main

import (
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
	doGreetManyTimes(c)
}
