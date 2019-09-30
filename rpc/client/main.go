package main

import (
	pb "github.com/buffge/gobyexample/rpc/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address = ":50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)
	r, err := c.GetUser(context.Background(), &pb.UserQueryInfo{Id: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%v", r)
}
