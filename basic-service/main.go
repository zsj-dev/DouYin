package main

import (
	"log"
	"net"

	"github.com/zsj-dev/DouYin/basic-service/controller"
	"github.com/zsj-dev/DouYin/basic-service/initialization"
	"github.com/zsj-dev/DouYin/pb"
	"google.golang.org/grpc"
)

func main() {
	initialization.RegisterMySQL()
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &controller.UserServerImpl{})
	pb.RegisterPublishServiceServer(server, &controller.PublishServerImpl{})
	pb.RegisterFeedServiceServer(server, &controller.FeedServiceImpl{})
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("user service init error: %v", err)
	}
	server.Serve(listen)
}
