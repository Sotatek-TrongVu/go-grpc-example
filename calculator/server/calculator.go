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

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	number := in.GetNumber()
	var k int32 = 2
	for number > 1 {
		if number%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Prime: k,
			})
			number = number / k
		} else {
			k = k + 1
		}
	}
	return nil
}
