# TCP 聊天室
```bash
# server
go run cmd/tcp/server/main.go
# client
go run cmd/tcp/client/main.go
go run cmd/tcp/client/main.go
go run cmd/tcp/client/main.go
```

# WebSocket 聊天室
```bash
# server
go run cmd/websocket/nhooyr_server/main.go   # nhooyr.io/websocket 版本
#   or
go run cmd/websocket/gorilla_server/main.go  # gorilla/websocket 版本

# client
go run cmd/websocket/client/main.go 
```

# 完整的聊天室 Demo
```bash
go run cmd/chatroom/main.go
```

# 性能测试
```bash
# 启动聊天室
go build -o chatroom cmd/chatroom/main.go
./chatroom
# 尝试 10 个用户同时进入聊天室, 并每 20s 各发送一条消息
go run cmd/benchmark/main.go -u 10 -m 20s -l 0
go run cmd/benchmark/main.go -u 200 -m 20s -l 10ms
```