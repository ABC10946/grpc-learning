package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"simple/gen/api"

	"google.golang.org/grpc"
)

type server struct {
	api.UnimplementedSimpleSearchServiceServer
}

var (
	data = make(map[string]api.Person)
)

func init() {
	data["john"] = api.Person{Name: "john", Age: 20, Email: "john@local"}
	data["jane"] = api.Person{Name: "jane", Age: 20, Email: "jane@local"}
	data["doe"] = api.Person{Name: "doe", Age: 20, Email: "doe@local"}
	log.Printf("Data: %v", data)
}

func (s *server) Search(ctx context.Context, req *api.SearchRequest) (*api.SearchResponse, error) {
	query := req.GetQuery()
	log.Printf("Query: %s", query)
	if data[query].Name != "" {
		person := data[query]
		return &api.SearchResponse{Person: &person, IsFound: true}, nil
	}

	return &api.SearchResponse{IsFound: false}, nil
}

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterSimpleSearchServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
