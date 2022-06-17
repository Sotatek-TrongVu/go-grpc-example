package main

import (
	"context"
	"fmt"

	pb "github.com/hiiamtrong/go-grpc-example/greet/proto"
)

func (s *Server) SayHello(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Printf("hehe\n")
	name := req.GetName()

	return &pb.GreetResponse{
		Result: "hello " + name,
	}, nil
}
