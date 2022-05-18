package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"grpc_test/Helloworld/proto"
)

func main() {
	md := metadata.New(map[string]string{"name": "lcq"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	conn, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	defer conn.Close()
	client := proto.NewHelloworldClient(conn)
	var header metadata.MD
	res, err := client.Sayhello(ctx, &proto.HelloReq{Name: "zrf"}, grpc.Header(&header))
	fmt.Printf("获取的元数据为%s\n", header)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
