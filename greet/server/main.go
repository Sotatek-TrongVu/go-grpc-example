package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/hiiamtrong/go-grpc-example/greet/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

func main() {
	addr := fmt.Sprintf(
		"localhost:%d", 50051,
	)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("RPC server listen on %s", addr)

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

}
