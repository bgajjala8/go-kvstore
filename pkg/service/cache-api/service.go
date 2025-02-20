package service

import (
	pb "github.com/bgajjala8/go-cache/proto-public/go"
)

// Service implements the gRPC service "CacheService".
type Service struct {
	pb.UnimplementedCacheServiceServer
	cache map[string]string
}

// NewService creates a new Service.
func NewService() *Service {
	return &Service{
		cache: make(map[string]string),
	}
}

// Ensure that Service implements the CacheServiceServer interface.
var (
	_ pb.CacheServiceServer = (*Service)(nil)
)
