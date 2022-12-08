package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "go-grpc/proto"
)

var port = flag.Int("port", 50051, "The server port")

type server struct {
	pb.UnimplementedItemsServer
}

func (s *server) Create(ctx context.Context, in *pb.RequestItem) (*pb.ResponseItem, error) {
	fmt.Printf("Generating item: %v\n", in.GetDescription())

	response := pb.ResponseItem{
		Description: in.GetDescription(),
		Time:        timestamppb.New(time.Now()),
		Val1:        9 * 3,
		Val2:        float32(9*2) / float32(.3),
		Opt:         pb.Option_OPTION_TWO,
	}

	return &response, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterItemsServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
