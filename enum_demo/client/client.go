package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_test/enum_demo/proto"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	client := proto.NewGreetClient(conn)
	res, _ := client.SayHello(context.Background(), &proto.Req{
		Name: "lcq",
		Sex:  proto.Gender_woman,
	})
	fmt.Println(res)
}
