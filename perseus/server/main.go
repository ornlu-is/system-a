package main

import (
	"fmt"
	"log"
	"net"

	perseus "github.com/ornlu-is/system-a/perseus"
	"google.golang.org/grpc"
)

const (
	servicePort = 50051
)

type server struct {
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", servicePort))
	if err != nil {
		log.Fatalf("failed to listen to port %d: %s", servicePort, err.Error())
	}

	srv := grpc.NewServer()

	perseus.RegisterPerseusServer(srv, &server{})

	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
