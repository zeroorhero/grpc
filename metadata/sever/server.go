package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"net"

	"grpc_test/Helloworld/proto"

	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) Sayhello(ctx context.Context, client *proto.HelloReq) (*proto.HelloRes, error) {
	data, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(data.Get("name"))
	md := metadata.Pairs("data", "zrf")
	_ = grpc.SendHeader(ctx, md)
	return &proto.HelloRes{Data: "lcq"}, nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterHelloworldServer(server, &Server{})
	listener, _ := net.Listen("tcp", "127.0.0.1:8080")
	err := server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
