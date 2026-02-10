package service

import (
	"context"
	"patijournal/internal/repository"
	"patijournal/pkg/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EntryService struct {
	pb.UnimplementedEntryServiceServer
	repo repository.EntryRepository
}

func NewEntryService(repo repository.EntryRepository) *EntryService {
	return &EntryService{
		repo: repo,
	}
}

func (s *EntryService) GetEntries(ctx context.Context, req *pb.GetEntriesRequest) (*pb.GetEntriesResponse, error) {
	entries, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to retrieve entries: %v", err)
	}

	return &pb.GetEntriesResponse{
		Entries: entries,
	}, nil
}

func (s *EntryService) GetEntry(ctx context.Context, req *pb.GetEntryRequest) (*pb.GetEntryResponse, error) {
	if req.Id < 0 {
		return nil, status.Error(codes.InvalidArgument, "entry ID is required")
	}

	entry, err := s.repo.GetByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "entry not found: %v", err)
	}

	return &pb.GetEntryResponse{
		Entry: entry,
	}, nil
}

func (s *EntryService) CreateEntry(ctx context.Context, req *pb.CreateEntryRequest) (*pb.CreateEntryResponse, error) {
	return nil, status.Errorf(codes.Unknown, "TODO: Implement create")
}

func (s *EntryService) UpdateEntry(ctx context.Context, req *pb.UpdateEntryRequest) (*pb.UpdateEntryResponse, error) {
	return nil, status.Errorf(codes.Unknown, "TODO: Implement update")
}
