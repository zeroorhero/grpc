package main

import (
	"context"
	"fmt"
	"net"

	"grpc_test/Helloworld/proto"

	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) Sayhello(context.Context, *proto.HelloReq) (*proto.HelloRes, error) {
	return &proto.HelloRes{Data: "lcq"}, nil
}

func main() {
	inter := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("请求来了")
		res, err := handler(ctx, req)
		fmt.Println("请求结束了")
		return res, err
	}
	opt := grpc.UnaryInterceptor(inter)
	server := grpc.NewServer(opt)
	proto.RegisterHelloworldServer(server, &Server{})
	listener, _ := net.Listen("tcp", "127.0.0.1:8080")
	err := server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
