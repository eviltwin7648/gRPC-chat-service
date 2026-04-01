package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/eviltwin7648/gRPC-chat-service/gen/chat"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	stream, err := client.ChatStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter username: ")
	scanner.Scan()
	username := scanner.Text()
	fmt.Print("Enter RoomID: ")
	scanner.Scan()
	roomId := scanner.Text()

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("[%s]: %s\n", msg.User, msg.Message)
		}
	}()

	for scanner.Scan() {
		text := scanner.Text()
		err := stream.Send(&pb.ChatMessage{
			User:      username,
			Message:   text,
			RoomId:    roomId,
			Timestamp: time.Now().Unix(),
		})
		if err != nil {
			log.Println("send error:", err)
			break
		}
	}
}
