package service

import (
	"context"

	pb "github.com/bgajjala8/go-cache/proto-public/go"
)

func (s *Service) Delete(context.Context, *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return nil, nil
}
