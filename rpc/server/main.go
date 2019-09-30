package main

import (
	"context"
	"errors"
	pb "github.com/buffge/gobyexample/rpc/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type User struct {
	Name string
	age  int32
}
type server struct{}

var users map[int32]User

func (s *server) GetUser(ctx context.Context, in *pb.UserQueryInfo) (*pb.UserInfo, error) {
	uid := in.Id
	if _, ok := users[uid]; !ok {
		return nil, errors.New("用户不存在")
	}
	log.Println("正在输出用户信息", users[uid])
	return &pb.UserInfo{Name: users[uid].Name, Age: users[uid].age}, nil
}

func main() {
	users = make(map[int32]User)
	users[1] = User{
		"buffge",
		25,
	}
	users[2] = User{
		"zty",
		24,
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	s.Serve(lis)
}
