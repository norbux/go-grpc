package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "go-grpc/proto"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewItemsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Create(ctx, &pb.RequestItem{Description: "Test"})
	if err != nil {
		log.Fatalf("failed to create new item: %v", err)
	}

	fmt.Println("Item:")
	fmt.Printf("Description: %s\n", response.Description)
	fmt.Printf("Time: %v\n", response.Time.AsTime())
	fmt.Printf("Val1: %d\n", response.Val1)
	fmt.Printf("Val2: %f\n", response.Val2)
	fmt.Printf("Option: %v\n", response.Opt)
}
