package main

import (
	"context"
	"flag"
	"log"

	"simple/gen/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	name = flag.String("name", "john", "name to search")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()

	c := api.NewSimpleSearchServiceClient(conn)
	r, err := c.Search(context.Background(), &api.SearchRequest{Query: *name})
	if err != nil {
		log.Fatalf("failed to search: %v", err)
	}

	log.Printf("Person: %v", r.GetPerson())

}
