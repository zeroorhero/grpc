package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_test/streamrpc/proto"
	"net"
	"sync"
	"time"
)

type Server struct{}

func (s *Server) Getdata(req *proto.Req, dataServer proto.Greet_GetdataServer) error {
	i := 0
	for {
		if i >= 10 {
			break
		}
		dataServer.Send(&proto.Res{Data: fmt.Sprintf("%s : %d", req.Name, i)})
		i++
		time.Sleep(time.Second)
	}
	return nil
}
func (s *Server) Putdata(dataServer proto.Greet_PutdataServer) error {
	for {
		res, err := dataServer.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(res)
	}
	return nil
}
func (s *Server) Alldata(dataServer proto.Greet_AlldataServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			err := dataServer.Send(&proto.Res{Data: "server"})
			time.Sleep(time.Second)
			if err != nil {
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			res, err := dataServer.Recv()
			if err != nil {
				return
			}
			fmt.Println(res)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterGreetServer(server, &Server{})
	listener, _ := net.Listen("tcp", "127.0.0.1:8080")
	err := server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
