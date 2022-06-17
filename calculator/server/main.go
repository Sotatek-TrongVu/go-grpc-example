package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/hiiamtrong/go-grpc-example/calculator/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.CalculatorServiceServer
}

func main() {

	addr := fmt.Sprintf(
		"localhost:%d", 50052)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer lis.Close()

	log.Printf("RPC server listen on %s", addr)

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
