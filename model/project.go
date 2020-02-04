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
			ID:   feature.Id,
			User: feature.User,
			IP:   feature.Ip,
			Map:  feature.Map,
		})

	}
	return results, nil
}

// CreateProject create project
func CreateProject(project *Project) (bool, string) {
	client := pb.NewMongoDBClient(MongoGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := client.InsertProject(ctx, &pb.ProjectInformation{
		Id:   project.ID,
		User: project.User,
		Ip:   project.IP,
		Map:  project.Map,
	})
	return result.IsVaild, result.Value
}

// UpdateProject update project
func UpdateProject(project *Project) (bool, string) {
	client := pb.NewMongoDBClient(MongoGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := client.UpdateProject(ctx, &pb.ProjectInformation{
		Id:   project.ID,
		User: project.User,
		Ip:   project.IP,
		Map:  project.Map,
	})
	return result.IsVaild, result.Value
}

// DeleteProject delete project
func DeleteProject(id string) (bool, string) {
	client := pb.NewMongoDBClient(MongoGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, _ := client.DeleteProject(ctx, &pb.ProjectId{
		Id: id,
	})
	return result.IsVaild, result.Value
}
