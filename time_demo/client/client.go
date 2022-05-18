package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc_test/time_demo/proto"
	"time"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	client := proto.NewGreetClient(conn)
	res, _ := client.SayHello(context.Background(), &proto.Req{
		Name: "lcq",
		Now:  timestamppb.New(time.Now()),
	})
	fmt.Println(res.Data)
}
