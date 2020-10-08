package gcloud

import (
	"context"

	"cloud.google.com/go/storage"
)

// Storage is a wrapper for Google Cloud storage client
type Storage struct {
	*storage.Client
}

// Init creates new cloud storage client
func Init(ctx context.Context) (*Storage, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &Storage{client}, nil
}
