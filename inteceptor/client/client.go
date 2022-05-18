package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_test/Helloworld/proto"
	"time"
)

func main() {
	inter := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc)
		fmt.Println(time.Since(start))
		return err
	}
	opt := grpc.WithUnaryInterceptor(inter)
	conn, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), opt)
	defer conn.Close()
	client := proto.NewHelloworldClient(conn)
	res, err := client.Sayhello(context.Background(), &proto.HelloReq{Name: "zrf"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
