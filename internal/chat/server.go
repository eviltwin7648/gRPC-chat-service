package chat

import (
	"log"

	pb "github.com/eviltwin7648/gRPC-chat-service/gen/chat"
	"github.com/google/uuid"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
	hub *Hub
}

func NewChatServer(hub *Hub) *ChatServer {
	return &ChatServer{hub: hub}
}

func (s *ChatServer) ChatStream(stream pb.ChatService_ChatStreamServer) error {
	clientId := uuid.New().String()
	var roomId string
	defer func() {
		if roomId != "" {
			s.hub.RemoveClient(roomId, clientId)
		}
	}()
	for {
		msg, err := stream.Recv()
		log.Printf("Message received", msg)
		if err != nil {
			return nil
		}
		if roomId == "" {
			roomId = msg.RoomId
			s.hub.AddClient(roomId, clientId, stream)
			continue
		}
		s.hub.Broadcast(roomId, msg)
	}
}
