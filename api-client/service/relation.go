package service

import (
	"log"

	"github.com/zsj-dev/DouYin/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var RelationClient pb.RelationServiceClient

func RelationConn() {
	conn, err := grpc.Dial(":50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println(err)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	RelationClient = pb.NewRelationServiceClient(conn)

}
