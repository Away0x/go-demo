package main

import (
	"context"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := GetClientConn(context.Background(), "localhost:8081", []grpc.DialOption{
		grpc.WithBlock(),
	})
	defer conn.Close()
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
