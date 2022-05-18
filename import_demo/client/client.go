package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc_test/import_demo/proto"
)

func main() {
	con, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	c := proto.NewGreetClient(con)
	res, _ := c.SayHello(context.Background(), &emptypb.Empty{})
	fmt.Println(res.Name)
}
