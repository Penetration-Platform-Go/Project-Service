package model

import (
	"log"
	"os"

	"google.golang.org/grpc"
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
	ID    string `json:"id,omitempty"`
	User  string `json:"user,omitempty"`
	Value string `json:"value,omitempty"`
}
