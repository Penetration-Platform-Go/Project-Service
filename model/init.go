package model

import (
	pb "github.com/Penetration-Platform-Go/gRPC-Files/MongoDB-Service"
	"google.golang.org/grpc"
	"log"
	"os"
)

// MongoGrpcClient for connection auth grpc service
var MongoGrpcClient *grpc.ClientConn

func init() {
	// get user service address
	AUTHADDRESS := "localhost:8083"
	conn, err := grpc.Dial(AUTHADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	MongoGrpcClient = conn
}

// Project define
type Project struct {
	ID        string          `json:"id,omitempty"`
	User      string          `json:"user,omitempty"`
	Title     string          `json:"title,omitepty"`
	Score     int32           `json:"score,omitepty"`
	Equipment []*pb.Equipment `json:"equipment,omitempty"`
	Map       *pb.Map         `json:"map,omitempty"`
}
