package main

import (
	"context"
	"fmt"
	"io"
	"log"

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

func (s *Server) SayLongTime(stream pb.GreetService_SayLongTimeServer) error {

	result := ""
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		result += in.GetName() + "\n"
	}

	stream.SendAndClose(&pb.GreetResponse{
		Result: result,
	})

	return nil
}

func (s *Server) SayEveryOne(stream pb.GreetService_SayEveryOneServer) error {
	log.Println("SayEveryOne is invoked")
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("failed to receive a request: %v", err)
		}
		log.Printf("Received a request: %v", in.GetName())

		err = stream.Send(&pb.GreetResponse{
			Result: "hello " + in.GetName(),
		})
		if err != nil {
			log.Fatalf("failed to send a response: %v", err)
		}
	}

}
