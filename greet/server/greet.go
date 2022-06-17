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

func (s *Server) SayManyTimes(in *pb.GreetRequest, stream pb.GreetService_SayManyTimesServer) error {
	for i := 0; i < 10; i++ {
		stream.Send(&pb.GreetResponse{
			Result: fmt.Sprintf("%d hello "+in.GetName(), i),
		})
		fmt.Println(i)
	}

	return nil
}
