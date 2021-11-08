```bash
# 安装 protoc 编译器
wget https://github.com/google/protobuf/releases/download/v3.11.2/protobuf-all-3.11.2.zip
unzip protobuf-all-3.11.2.zip && cd protobuf-3.11.2/
./configure
make && make install

protoc --version # 如报错 while loading shard libraries, 需要先执行 ldconfig 命令

# 安装 proto 的 go 语言插件
go install github.com/golang/protobuf/protoc-gen-go@v1.3.3
# 安装 grpc 插件
go install google.golang.org/grpc@v1.29.1
# 调试 grpc
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
:'
需要 grpc server 注册放射服务
import "google.golang.org/grpc/reflection"
// ...
s := grpc.NewServer()
reflection.Register(s)
'
grpcurl -plaintext localhost:8081 list
grpcurl -plaintext localhost:8081 list proto.TagService
grpcurl -plaintext -d '{"name": "go"}' localhost:8081 proto.TagService.GetTagList
```

# Unary RPC
> 一元 RPC，也称为单次 RPC，简单来讲，就是客户端发起一次普通的 RPC 请求，是最基础的调用，也是最常用的方式

```proto
rpc SayHello (HelloRequest) returns (HelloReply) {}
```
```go
// server
package main

import (
	"context"
	pb "learngrpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello world"}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, err := net.Listen("tcp", ":8081")
	if err != nil { log.Fatal(err) }
	if err = server.Serve(lis); err != nil { log.Fatal(err) }
}
```
```go
// client
package main

import (
	"context"
	pb "learngrpc/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial(":8081", grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_ = SayHello(client)
}

func SayHello(client pb.GreeterClient) error {
    // 发送 rpc 请求
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "demo"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}
// print: "client.SayHello resp: hello world"
```

# Server-side streaming RPC
> 服务端流式RPC是一个单向流，指Server为Stream、Client为普通的一元RPC请求, 简单来讲，就是客户端发起一次普通的RPC请求，服务端通过流式响应多次发送数据集，客户端Recv接收数据集

```proto
rpc SayList (HelloRequest) returns (stream HelloReply) {}
```
```go
// server
package main

import (
	"fmt"
	pb "learngrpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GreeterServer struct{}

func (s *GreeterServer) SayList(r *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	for n := 0; n <= 6; n++ {
		_ = stream.Send(&pb.HelloReply{Message: []string{fmt.Sprintf("hello world %d", n)}})
	}
	return nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
```
```go
// client
package main

import (
	"context"
	"io"
	pb "learngrpc/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial(":8081", grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_ = SayList(client, &pb.HelloRequest{Name: "demo"})
}

func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayList(context.Background(), r)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v", resp)
	}

	return nil
}
/**
2021/10/18 11:02:18 resp: message:"hello world 0" 
2021/10/18 11:02:18 resp: message:"hello world 1" 
2021/10/18 11:02:18 resp: message:"hello world 2" 
2021/10/18 11:02:18 resp: message:"hello world 3" 
2021/10/18 11:02:18 resp: message:"hello world 4" 
2021/10/18 11:02:18 resp: message:"hello world 5" 
2021/10/18 11:02:18 resp: message:"hello world 6" 
*/
```

# Client-side streaming RPC
> 客户端流式RPC是一个单向流，客户端通过流式发起多次RPC请求给服务端，而服务端仅发起一次响应给客户端

```proto
rpc SayRecord(stream HelloRequest) returns (HelloReply) {}
```
```go
// server
package main

import (
	"io"
	pb "learngrpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GreeterServer struct{}

func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			message := &pb.HelloReply{Message: []string{"say record"}}
			return stream.SendAndClose(message)
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v", resp)
	}
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
/**
2021/10/18 11:29:14 resp: name:"demo" 
2021/10/18 11:29:14 resp: name:"demo" 
2021/10/18 11:29:14 resp: name:"demo" 
2021/10/18 11:29:14 resp: name:"demo" 
2021/10/18 11:29:14 resp: name:"demo" 
2021/10/18 11:29:14 resp: name:"demo"
*/
```
```go
// client
package main

import (
	"context"
	pb "learngrpc/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial(":8081", grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_ = SayRecord(client, &pb.HelloRequest{Name: "demo"})
}

func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRecord(context.Background())
	for n := 0; n < 6; n++ {
		_ = stream.Send(r)
	}
	resp, _ := stream.CloseAndRecv()

	log.Printf("resp: %v", resp)
	return nil
}
// 2021/10/18 11:31:19 resp: message:"say record"
```

# Bidirectional streaming RPC
双向流式RPC，顾名思义是双向流，由客户端以流式的方式发起请求，服务端同样以流式的方式响应请求

首个请求一定是由客户端发起的，但具体的交互方式（谁先谁后、一次发多少、响应多少、什么时候关闭）则由程序编写的方式来确定（可以结合协程）

```proto
rpc SayRoute(stream HelloRequest) returns (stream HelloReply) {}
```
```go
// server
package main

import (
	"fmt"
	"io"
	pb "learngrpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GreeterServer struct{}

func (s *GreeterServer) SayRoute(stream pb.Greeter_SayRouteServer) error {
	n := 0
	for {
		_ = stream.Send(&pb.HelloReply{Message: []string{fmt.Sprintf("say route %d", n)}})

		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++
		log.Printf("resp: %v", resp)
	}
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
/**
2021/10/18 11:35:39 resp: name:"demo" 
2021/10/18 11:35:39 resp: name:"demo" 
2021/10/18 11:35:39 resp: name:"demo" 
2021/10/18 11:35:39 resp: name:"demo" 
2021/10/18 11:35:39 resp: name:"demo" 
2021/10/18 11:35:39 resp: name:"demo" 
*/
```
```go
// client
package main

import (
	"context"
	"io"
	pb "learngrpc/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial(":8081", grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_ = SayRoute(client, &pb.HelloRequest{Name: "demo"})
}

func SayRoute(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRoute(context.Background())
	for n := 0; n < 6; n++ {
		_ = stream.Send(r)
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v", resp)
	}

	_ = stream.CloseSend()
	return nil
}
/**
2021/10/18 11:35:39 resp: message:"say route 0" 
2021/10/18 11:35:39 resp: message:"say route 1" 
2021/10/18 11:35:39 resp: message:"say route 2" 
2021/10/18 11:35:39 resp: message:"say route 3" 
2021/10/18 11:35:39 resp: message:"say route 4" 
2021/10/18 11:35:39 resp: message:"say route 5"
*/
```