package main

import (
	"context"
	"net"

	"grpc_test/Helloworld/proto"

	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) Sayhello(context.Context, *proto.HelloReq) (*proto.HelloRes, error) {
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
