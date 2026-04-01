# gRPC Chat Service (Go)

A minimal real-time chat system built using **gRPC bidirectional streaming in Go**.
Supports multiple clients and room-based message routing using an in-memory hub.

---

## 🚀 Features

* Bidirectional streaming using gRPC
* CLI-based clients
* Multiple clients support
* Room-based messaging (1–1 and group chats unified)
* Concurrent message broadcasting
* Clean separation of concerns (server / hub / client)

---

## 🧠 Architecture Overview

```
Client (CLI)
   ↕ gRPC Stream
Chat Server
   ↕
Hub (in-memory)
   ↕
Room → Clients map
```

---

## 📦 Project Structure

```
cmd/
  ├── server/        # Entry point for gRPC server
  └── client/        # CLI client

internal/chat/
  ├── hub.go         # Core logic (rooms, clients, broadcast)
  ├── server.go      # gRPC service implementation

gen/chat/            # Generated protobuf code

proto/
  └── chat.proto     # gRPC contract
```

---

## 🔧 How It Works

### 1. Client Flow

* Connects to server using gRPC
* Opens a bidirectional stream
* Sends first message containing `room_id` (used as join)
* Sends messages via stdin
* Receives messages concurrently

---

### 2. Server Flow

```
Client connects
   ↓
Stream opened
   ↓
First message → register client to room
   ↓
Subsequent messages → broadcast to room
```

---

### 3. Hub Logic

* Maintains:

  ```
  map[roomID]map[clientID]stream
  ```
* Handles:

  * client registration
  * client removal
  * broadcasting messages to all clients in a room

---

## 🧪 Running the Project

### 1. Generate protobuf code

```
protoc \
  --go_out=. --go-grpc_out=. \
  proto/chat.proto
```

---

### 2. Run server

```
go run cmd/server/main.go
```

---

### 3. Run client(s)

```
go run cmd/client/main.go
```

---

### 4. Example

Terminal 1:

```
username: user1
room: room1
hello
```

Terminal 2:

```
username: user2
room: room1
hi
```

👉 Both users will receive each other's messages.

---

## ⚠️ Limitations (Current Version)

* Single-node (no horizontal scaling)
* No message persistence
* No authentication
* No ordering guarantees
* No backpressure handling
* In-memory state only

---

## 📌 Design Decisions

* Used **gRPC streaming** instead of WebSockets for simplicity and performance
* Unified **1–1 and group chat** using room abstraction
* Used **in-memory hub** for fast routing
* Used **goroutines for non-blocking broadcast**

---

## 🧭 Future Improvements

* Replace hub with Redis Pub/Sub for multi-node scaling
* Add message persistence (Postgres / Kafka)
* Handle slow clients (backpressure)
* Add authentication (JWT)
* Add delivery acknowledgements

---

## 🎯 Learning Outcomes

* Understanding gRPC bidirectional streaming
* Managing concurrent client connections
* Designing in-memory routing systems
* Handling real-time communication patterns

---

## 🧑‍💻 Author

Built as a focused learning project to understand real-time systems using Go and gRPC.
