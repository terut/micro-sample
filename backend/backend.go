package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/terut/micro-sample/proto/gen"
	context "golang.org/x/net/context"
)

func main() {
	port := 8000
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to create logger :%v", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}
	server := grpc.NewServer(grpc.UnaryInterceptor(
		grpc_zap.UnaryServerInterceptor(logger),
	))
	grpc_zap.ReplaceGrpcLogger(logger)
	pb.RegisterWordServer(server, &WordService{})
	server.Serve(lis)
}

type WordService struct{}

func (s *WordService) Build(ctx context.Context, req *pb.WordRequest) (*pb.WordResponse, error) {
	val := req.Val
	return &pb.WordResponse{Val: fmt.Sprintf("Hello,%s", val)}, nil
}
