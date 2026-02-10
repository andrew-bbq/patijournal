package repository

import (
	"context"
	"errors"
	"patijournal/pkg/pb"
)

type EntryRepository interface {
	GetAll(ctx context.Context) ([]*pb.Entry, error)
	GetByID(ctx context.Context, id int32) (*pb.Entry, error)
	Create(ctx context.Context, entry *pb.Entry) (*pb.Entry, error)
	Update(ctx context.Context, entry *pb.Entry) (*pb.Entry, error)
}

type InMemoryRepository struct {
	entries []*pb.Entry
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		entries: []*pb.Entry{
			{
				Id:    1,
				Title: "First Entry",
				Body:  "First entry body",
				Image: "data:image/gif;base64,R0lGODlhAQABAAAAACw=",
			},
			{
				Id:    2,
				Title: "Second Entry",
				Body:  "Second entry body",
				Image: "data:image/gif;base64,R0lGODlhAQABAAAAACw=",
			},
		},
	}
}

func (r *InMemoryRepository) GetAll(ctx context.Context) ([]*pb.Entry, error) {
	return r.entries, nil
}

func (r *InMemoryRepository) GetByID(ctx context.Context, id int32) (*pb.Entry, error) {
	for _, entry := range r.entries {
		if entry.Id == id {
			return entry, nil
		}
	}
	return nil, errors.New("Entry not found")
}

func (r *InMemoryRepository) Create(ctx context.Context, entry *pb.Entry) (*pb.Entry, error) {
	entry.Id = int32(len(r.entries) + 1)
	r.entries = append(r.entries, entry)
	return entry, nil
}

func (r *InMemoryRepository) Update(ctx context.Context, entry *pb.Entry) (*pb.Entry, error) {
	for i, e := range r.entries {
		if e.Id == entry.Id {
			r.entries[i] = entry
			return entry, nil
		}
	}
	return nil, errors.New("Entry not found")
}
