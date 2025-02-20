package service

import (
	"context"

	pb "github.com/bgajjala8/go-cache/proto-public/go"
)

func (s *Service) Set(context.Context, *pb.SetRequest) (*pb.SetResponse, error) {
	return nil, nil
}
