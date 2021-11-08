package main

import (
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "learngrpc/proto"
	"learngrpc/server"
)

const httpPort = "8081"
const rpcPort = "8082"

func main() {
	errs := make(chan error)

	go func() {
		err := runHTTPServer(httpPort)
		if err != nil {
			errs <- err
		}
	}()

	go func() {
		err := runGrpcServer(rpcPort)
		if err != nil {
			errs <- err
		}
	}()

	err := <-errs
	log.Fatalf("Run Server err: %v", err)
}

func runHTTPServer(port string) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		_, _ = rw.Write([]byte("pong"))
	})

	return http.ListenAndServe(":"+port, serveMux)
}

func runGrpcServer(port string) error {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	return s.Serve(lis)
}
