package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_test/Helloworld/proto"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	defer conn.Close()
	client := proto.NewHelloworldClient(conn)
	res, err := client.Sayhello(context.Background(), &proto.HelloReq{Name: "zrf"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
