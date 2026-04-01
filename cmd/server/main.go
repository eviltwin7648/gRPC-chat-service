package main

import (
	"log"
	"net"

	pb "github.com/eviltwin7648/gRPC-chat-service/gen/chat"
	chat "github.com/eviltwin7648/gRPC-chat-service/internal/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	hub := chat.NewHub()
	server := chat.NewChatServer(hub)

	pb.RegisterChatServiceServer(grpcServer, server)

	log.Println("Server Running on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
