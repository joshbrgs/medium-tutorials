package nemesis

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/medium-tutorials/bad-inc/cmd/nemesis/api/gen"

	"google.golang.org/grpc"
)

type NemesisServiceServer struct {
	pb.UnimplementedNemesisServiceServer
	nemesisStore map[string]*pb.NemesisResponse
}

func NewNemesisServiceServer() *NemesisServiceServer {
	return &NemesisServiceServer{
		nemesisStore: make(map[string]*pb.NemesisResponse),
	}
}

func (s *NemesisServiceServer) CreateNemesis(ctx context.Context, req *pb.CreateNemesisRequest) (*pb.NemesisResponse, error) {
	nemesisID := fmt.Sprintf("nemesis-%s", req.PersonId)
	nemesis := &pb.NemesisResponse{
		NemesisId:    nemesisID,
		PersonId:     req.PersonId,
		NemesisName:  req.NemesisName,
		NemesisPower: req.NemesisPower,
	}
	s.nemesisStore[nemesisID] = nemesis
	return nemesis, nil
}

func (s *NemesisServiceServer) GetNemesis(ctx context.Context, req *pb.GetNemesisRequest) (*pb.NemesisResponse, error) {
	nemesis, exists := s.nemesisStore[req.NemesisId]
	if !exists {
		return nil, fmt.Errorf("nemesis with ID %s not found", req.NemesisId)
	}
	return nemesis, nil
}

func (s *NemesisServiceServer) UpdateNemesis(ctx context.Context, req *pb.UpdateNemesisRequest) (*pb.NemesisResponse, error) {
	nemesis, exists := s.nemesisStore[req.NemesisId]
	if !exists {
		return nil, fmt.Errorf("nemesis with ID %s not found", req.NemesisId)
	}
	nemesis.NemesisName = req.NemesisName
	nemesis.NemesisPower = req.NemesisPower
	s.nemesisStore[req.NemesisId] = nemesis
	return nemesis, nil
}

func (s *NemesisServiceServer) DeleteNemesis(ctx context.Context, req *pb.DeleteNemesisRequest) (*pb.DeleteNemesisResponse, error) {
	_, exists := s.nemesisStore[req.NemesisId]
	if !exists {
		return nil, fmt.Errorf("nemesis with ID %s not found", req.NemesisId)
	}
	delete(s.nemesisStore, req.NemesisId)
	return &pb.DeleteNemesisResponse{Message: "Nemesis deleted successfully"}, nil
}

func (s *NemesisServiceServer) ListNemeses(ctx context.Context, req *pb.Empty) (*pb.ListNemesesResponse, error) {
	nemeses := make([]*pb.NemesisResponse, 0, len(s.nemesisStore))
	for _, nemesis := range s.nemesisStore {
		nemeses = append(nemeses, nemesis)
	}
	return &pb.ListNemesesResponse{Nemeses: nemeses}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterNemesisServiceServer(server, NewNemesisServiceServer())

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Nemesis Service is running on port 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
