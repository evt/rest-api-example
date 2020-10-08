package gcloud

import (
	"context"

	"cloud.google.com/go/storage"
)

// GCloudStorage is a wrapper for Google Cloud storage client
type GCloudStorage struct {
	*storage.Client
}

// Init creates new cloud storage client
func Init(ctx context.Context) (*GCloudStorage, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &GCloudStorage{client}, nil
}
