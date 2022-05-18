package main

import (
	"context"
	"google.golang.org/grpc"
	"net"

	"grpc_test/Array/proto"
)

type Server struct{}

func (s *Server) SayHello(c context.Context, req *proto.Req) (*proto.Res, error) {
	return &proto.Res{
		Data: "lcq",
		Id:   []int32{2, 2, 2, 2},
	}, nil
}

func main() {
	s := grpc.NewServer()
	proto.RegisterGreetServer(s, &Server{})
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	err = s.Serve(listener)
	if err != nil {
		panic(err)
	}
}
