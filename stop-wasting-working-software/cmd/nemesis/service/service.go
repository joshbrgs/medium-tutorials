package service

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/medium-tutorials/bad-inc/cmd/nemesis/api/gen"
	"github.com/medium-tutorials/bad-inc/cmd/nemesis/models"
	"github.com/medium-tutorials/bad-inc/pkgs/server"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NemesisServiceServer struct {
	pb.UnimplementedNemesisServiceServer
	db *mongo.Database
	mq *server.RabbitMQ
}

func NewNemesisServiceServer(db *mongo.Database, mq *server.RabbitMQ) *NemesisServiceServer {
	return &NemesisServiceServer{db: db, mq: mq}
}

func (s *NemesisServiceServer) CreateNemesis(ctx context.Context, req *pb.CreateNemesisRequest) (*pb.NemesisResponse, error) {
log.Printf("Received CreateNemesisRequest: %+v", req)
	nemesis := &models.Nemesis{
		Name:  req.Name,
		Power: req.Power,
	}

	log.Printf("Nemesis: %s, %s", nemesis.Name, nemesis.Power)

	if err := nemesis.Create(ctx, s.db, "nemesis", nemesis); err != nil {
		log.Printf("Failed to create nemesis: %v", err)
		return nil, err
	}

	if err := s.mq.Publish(ctx, "nemesis", fmt.Sprintf("New Nemesis created: %s with %s powers", nemesis.Name, nemesis.Power)); err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	return &pb.NemesisResponse{
		Id:    nemesis.ID.Hex(),
		Name:  nemesis.Name,
		Power: nemesis.Power,
	}, nil
}

func (s *NemesisServiceServer) GetNemesis(ctx context.Context, req *pb.GetNemesisRequest) (*pb.NemesisResponse, error) {
	var nemesis models.Nemesis

	filter := bson.M{"_id": req.Id}                             // Create filter for fetching by ID
	err := nemesis.Read(ctx, s.db, "nemesis", filter, &nemesis) // Use mongorm.Read

	return &pb.NemesisResponse{
		Id:    nemesis.ID.Hex(),
		Name:  nemesis.Name,
		Power: nemesis.Power,
	}, err
}

func (s *NemesisServiceServer) UpdateNemesis(ctx context.Context, req *pb.UpdateNemesisRequest) (*pb.NemesisResponse, error) {
	var nemesis models.Nemesis
	id := req.Id

	// Update fields
	nemesis.Name = req.Name
	nemesis.Power = req.Power
	nemesis.UpdatedAt = time.Now() // Update timestamp

	filter := bson.M{"_id": id}       // Create filter for update
	update := bson.M{"$set": nemesis} // Update data

	err := nemesis.Update(ctx, s.db, "nemesis", filter, update) // Use mongorm.Update

	if err := s.mq.Publish(ctx, "nemesis", fmt.Sprintf("New Nemesis updated: %s with %s powers", nemesis.Name, nemesis.Power)); err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	return &pb.NemesisResponse{
		Id:    nemesis.ID.Hex(),
		Name:  nemesis.Name,
		Power: nemesis.Power,
	}, err
}

func (s *NemesisServiceServer) DeleteNemesis(ctx context.Context, req *pb.DeleteNemesisRequest) (*pb.Empty, error) {
	var nemesis models.Nemesis
	filter := bson.M{"_id": req.Id}
	// Create filter for deletion

	err := nemesis.Delete(ctx, s.db, "nemesis", filter) // Use mongorm.Delete

	if err := s.mq.Publish(ctx, "nemesis", fmt.Sprintf("Nemesis deleted: %s", req.Id)); err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	return &pb.Empty{}, err
}

func (s *NemesisServiceServer) ListNemesis(ctx context.Context, req *pb.Empty) (*pb.ListNemesisResponse, error) {
	var nemeses []models.Nemesis

	// Create a new Model instance to use the List method
	var nemesis models.Nemesis

	if err := nemesis.List(ctx, s.db, "nemesis", bson.M{}, &nemeses); err != nil {
		log.Printf("Failed to list nemeses: %v", err)
		return nil, err
	}

	var responses []*pb.NemesisResponse
	for _, nemesis := range nemeses {
		responses = append(responses, &pb.NemesisResponse{
			Id:    nemesis.ID.Hex(),
			Name:  nemesis.Name,
			Power: nemesis.Power,
		})
	}

	return &pb.ListNemesisResponse{Nemeses: responses}, nil
}
