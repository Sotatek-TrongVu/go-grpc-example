package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/hiiamtrong/go-grpc-example/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := fmt.Sprintf(
		"localhost:%d", 50052,
	)

	listen, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)

	}

	defer listen.Close()

	c := pb.NewCalculatorServiceClient(listen)

	rs, err := c.Sum(context.Background(), &pb.SumRequest{FirstNumber: 2, SecondNumber: 2})

	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}

	log.Println(rs.Result)
}
