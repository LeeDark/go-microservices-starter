package main

import (
	"context"
	"net"
	"testing"
	"time"

	pb "github.com/LeeDark/go-microservices-starter/grpc-playground/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

func TestGreeterUnaryRPCs(t *testing.T) {
	listener := bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})

	go func() {
		_ = grpcServer.Serve(listener)
	}()

	conn, err := grpc.NewClient(
		"passthrough:///bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("create gRPC client: %v", err)
	}
	t.Cleanup(func() {
		_ = conn.Close()
		grpcServer.Stop()
		_ = listener.Close()
	})

	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tests := []struct {
		name string
		call func(context.Context, *pb.HelloRequest) (*pb.HelloReply, error)
		want string
	}{
		{
			name: "SayHello",
			call: func(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
				return client.SayHello(ctx, request)
			},
			want: "Hello Lee",
		},
		{
			name: "SayHelloAgain",
			call: func(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
				return client.SayHelloAgain(ctx, request)
			},
			want: "Hello again Lee",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reply, err := tt.call(ctx, &pb.HelloRequest{Name: "Lee"})
			if err != nil {
				t.Fatalf("call RPC: %v", err)
			}
			if got := reply.GetMessage(); got != tt.want {
				t.Errorf("message = %q, want %q", got, tt.want)
			}
		})
	}
}
