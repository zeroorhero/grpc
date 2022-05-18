package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_test/insert_demo/proto"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	client := proto.NewGreetClient(conn)
	res, _ := client.SayHello(context.Background(), &proto.Req{Name: "zrf"})
	fmt.Println(res.Data)
	fmt.Println(res.In)
}
