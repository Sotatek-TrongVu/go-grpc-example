package main

import (
	"context"

	pb "github.com/hiiamtrong/go-grpc-example/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	first := req.GetFirstNumber()
	second := req.GetSecondNumber()

	return &pb.SumResponse{
		Result: first + second,
	}, nil
}
