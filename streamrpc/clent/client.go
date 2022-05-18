package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_test/streamrpc/proto"
	"sync"
	"time"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	defer conn.Close()
	client := proto.NewGreetClient(conn)
	res, _ := client.Getdata(context.Background(), &proto.Req{Name: "lcq"})
	for {
		mes, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(mes)
	}
	putdata, _ := client.Putdata(context.Background())
	i := 0
	for {
		if i >= 10 {
			break
		}
		err := putdata.Send(&proto.Req{Name: "zrf"})
		if err != nil {
			break
		}
		time.Sleep(time.Second)
		i++
	}
	alldata, _ := client.Alldata(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			err := alldata.Send(&proto.Req{Name: "client"})
			time.Sleep(time.Second)
			if err != nil {
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			res, err := alldata.Recv()
			if err != nil {
				return
			}
			fmt.Println(res)
		}
	}()
	wg.Wait()
}
