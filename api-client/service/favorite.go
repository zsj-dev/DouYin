package service

import (
	"log"

	"github.com/zsj-dev/DouYin/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var FavoriteClient pb.FavoriteServiceClient

func FavoriteConn() {
	conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(kacp))
	log.Println(err)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()
	FavoriteClient = pb.NewFavoriteServiceClient(conn)
}
