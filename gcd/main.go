package main

import (
	"context"
	"log"
	"net"

	"github.com/mramsden/kubernetes-go-grpc-tutorial/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGCDServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (s *server) Compute(ctx context.Context, r *pb.GCDRequest) (*pb.GCDResponse, error) {
	a, b := r.A, r.B
	for b != 0 {
		a, b = b, a%b
	}
	return &pb.GCDResponse{Result: a}, nil
}
