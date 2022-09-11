package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ornlu-is/system-a/perseus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error creating grpc connection")
	}
	perseusClient := perseus.NewPerseusClient(conn)

	ctx := context.Background()
	rsp, err := perseusClient.Get(ctx, &perseus.GetReq{
		Id: 123456,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp)
}
