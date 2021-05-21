package service

import (
	"context"
	"fmt"

	pb "github.com/buffge/gobyexample/third/kratos/api/helloworld"
	"github.com/go-kratos/kratos/v2/log"
)

type HelloworldService struct {
	pb.UnimplementedHelloworldServer
	log *log.Helper
}

func NewHelloworldService(logger log.Logger) *HelloworldService {
	return &HelloworldService{
		log: log.NewHelper("service/buffge", logger),
	}
}

func (s *HelloworldService) CreateHelloworld(ctx context.Context, req *pb.CreateHelloworldRequest) (*pb.CreateHelloworldReply, error) {
	return &pb.CreateHelloworldReply{}, nil
}
func (s *HelloworldService) UpdateHelloworld(ctx context.Context, req *pb.UpdateHelloworldRequest) (*pb.UpdateHelloworldReply, error) {
	return &pb.UpdateHelloworldReply{}, nil
}
func (s *HelloworldService) DeleteHelloworld(ctx context.Context, req *pb.DeleteHelloworldRequest) (*pb.DeleteHelloworldReply, error) {
	return &pb.DeleteHelloworldReply{}, nil
}
func (s *HelloworldService) GetHelloworld(ctx context.Context, req *pb.GetHelloworldRequest) (*pb.GetHelloworldReply, error) {
	return &pb.GetHelloworldReply{
		Msg: fmt.Sprintf("你好%s,%dage!", req.Name, req.Age),
	}, nil
}
func (s *HelloworldService) ListHelloworld(ctx context.Context, req *pb.ListHelloworldRequest) (*pb.ListHelloworldReply, error) {
	return &pb.ListHelloworldReply{}, nil
}
