package model

import (
	"context"
	"io"
	"time"

	pb "github.com/Penetration-Platform-Go/gRPC-Files/MongoDB-Service"
)

// QueryProject Query Project By username
func QueryProject(username string) ([]Project, error) {
	client := pb.NewMongoDBClient(MongoGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.QueryProjectsByUsername(ctx, &pb.Username{Username: username})
	if err != nil {
		return nil, err
	}
	var results []Project
	for {
		feature, err := stream.Recv()
		if err == io.EOF || feature == nil {
			break
		}
		results = append(results, Project{
			ID:    feature.GetId(),
			User:  feature.GetUser(),
			Value: feature.GetValue().GetTemp(),
		})

	}
	return results, nil
}
