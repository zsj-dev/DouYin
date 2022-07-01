package main

import (
	"log"
	"net"

	"github.com/zsj-dev/DouYin/pb"
	"github.com/zsj-dev/DouYin/relation-service/controller"
	"github.com/zsj-dev/DouYin/relation-service/initialization"
	"google.golang.org/grpc"
)

func main() {
	initialization.RegisterMySQL()
	server := grpc.NewServer()
	pb.RegisterRelationServiceServer(server, &controller.RelationServiceImpl{})

	listen, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("user service init error: %v", err)
	}
	server.Serve(listen)
}
