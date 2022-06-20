package main

import (
	"context"
	"io"

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

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	count := 0
	sum := 0

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		count++
		sum += int(in.GetNumber())

	}

	stream.SendAndClose(&pb.AvgResponse{
		Avg: float64(sum) / float64(count),
	})

	return nil

}

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	max := int32(0)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if in.GetNumber() <= max {
			continue
		}
		max = in.GetNumber()
		stream.Send(&pb.MaxResponse{
			Max: max,
		})
	}
	return nil
}
