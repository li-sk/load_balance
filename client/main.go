package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"load_balance/server/proto"
	"log"

	// 这里一定要import;很重要
	_ "github.com/mbobakov/grpc-consul-resolver"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(
		// consul://192.168.193.128:8500 consul地址
		// test-serve 拉取的服务名
		// wait=14s 等待时间
		// tag=manual 筛选条件
		"consul://192.168.193.128:8500/test-server?wait=14s&tag=manual",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := proto.NewTestClient(conn)
	for i := 0; i < 10; i++ {

		resp, err := client.Test(context.Background(), &emptypb.Empty{})
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.Msg)
	}
}
