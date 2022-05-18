package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_test/time_demo/proto"
	"net"
)

type Server struct{}

func (s *Server) SayHello(c context.Context, client *proto.Req) (*proto.Res, error) {
	return &proto.Res{Data: fmt.Sprintf("%d", client.Now.Seconds)}, nil
}

func main() {
	s := grpc.NewServer()
	proto.RegisterGreetServer(s, &Server{})
	listener, _ := net.Listen("tcp", "127.0.0.1:8080")
	_ = s.Serve(listener)
}
