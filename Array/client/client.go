package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"grpc_test/Array/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewGreetClient(conn)
	res, err := client.SayHello(context.Background(), &proto.Req{Name: "lcq"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Data)
	fmt.Println(res.Id)
}
