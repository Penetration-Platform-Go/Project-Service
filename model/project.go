package model

import (
	"context"
	"io"
	"time"

	pb "github.com/Penetration-Platform-Go/gRPC-Files/MongoDB-Service"
)

// QueryProjectsByUsername Query Project By username
func QueryProjectsByUsername(username string) ([]Project, error) {
	client := pb.NewMongoDBClient(MongoGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.QueryProjects(ctx, &pb.Condition{
		Value: []*pb.Value{
			{Key: "user", Value: username},
		},
	})
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
			ID:    feature.Id,
			User:  feature.User,
			Title: feature.Title,
			Score: feature.Score,
			IP:    feature.Ip,
			Map:   feature.Map,
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
		Id:    project.ID,
		User:  project.User,
		Title: project.Title,
		Ip:    project.IP,
		Map:   project.Map,
	})
	return result.IsVaild, result.Value
}

// UpdateProject update project
func UpdateProject(project *Project) (bool, string) {
	client := pb.NewMongoDBClient(MongoGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := client.UpdateProject(ctx, &pb.UpdateMessage{
		Condition: &pb.Condition{
			Value: []*pb.Value{
				{Key: "_id", Value: project.ID},
			},
		},
		Value: &pb.ProjectInformation{
			Ip:    project.IP,
			Map:   project.Map,
			Title: project.Title,
		},
		Key: []string{
			"ip", "map", "title",
		},
	})
	return result.IsVaild, result.Value
}

// DeleteProject delete project
func DeleteProject(id string) (bool, string) {
	client := pb.NewMongoDBClient(MongoGrpcClient)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, _ := client.DeleteProject(ctx, &pb.Condition{
		Value: []*pb.Value{
			{Key: "_id", Value: id},
		},
	})
	return result.IsVaild, result.Value
}
