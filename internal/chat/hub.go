package chat

import (
	"sync"

	pb "github.com/eviltwin7648/gRPC-chat-service/gen/chat"
)

type Hub struct {
	rooms map[string]map[string]pb.ChatService_ChatStreamServer // roomID : [(client id -> stream(live connection to client))]
	mu    sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		rooms: make(map[string]map[string]pb.ChatService_ChatStreamServer),
	}
}

func (h *Hub) AddClient(roomId, ClientId string, stream pb.ChatService_ChatStreamServer) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.rooms[roomId]; !ok {
		h.rooms[roomId] = make(map[string]pb.ChatService_ChatStreamServer)
	}
	h.rooms[roomId][ClientId] = stream
}

func (h *Hub) RemoveClient(roomId, ClientId string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if clients, ok := h.rooms[roomId]; ok {
		delete(clients, ClientId)
		if len(clients) == 0 {
			delete(h.rooms, roomId)
		}
	}

}

func (h *Hub) Broadcast(roomId string, msg *pb.ChatMessage) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, client := range h.rooms[roomId] {
		go client.Send(msg)
	}
}
