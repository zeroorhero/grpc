package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc_test/import_demo/proto"
	"net"
)

type Server struct{}

func (s *Server) SayHello(c context.Context, em *emptypb.Empty) (*proto.Pong, error) {
	return &proto.Pong{Name: "zrf"}, nil
}

func main() {
	s := grpc.NewServer()
	proto.RegisterGreetServer(s, &Server{})
	listener, _ := net.Listen("tcp", "127.0.0.1:8080")
	_ = s.Serve(listener)
}
