package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc_test/insert_demo/proto"
	"net"
)

type Server struct{}

func (s *Server) SayHello(c context.Context, client *proto.Req) (*proto.Res, error) {
	return &proto.Res{
		Data: "hhh",
		In:   []*proto.ResInsert{{Id: 2}, {Id: 666}},
	}, nil
}

func main() {
	s := grpc.NewServer()
	proto.RegisterGreetServer(s, &Server{})
	listener, _ := net.Listen("tcp", "127.0.0.1:8080")
	_ = s.Serve(listener)
}
