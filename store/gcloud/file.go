package gcloud

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/evt/rest-api-example/model"
	"github.com/pkg/errors"
)

// GoogleCloudFileContentRepo ...
type GoogleCloudFileContentRepo struct {
	storage *Storage
	bucket  string
}

// NewFileContentRepo ...
func NewFileContentRepo(storage *Storage, bucket string) *GoogleCloudFileContentRepo {
	return &GoogleCloudFileContentRepo{storage: storage, bucket: bucket}
}

// Upload file to Google Cloud storage
func (repo *GoogleCloudFileContentRepo) Upload(ctx context.Context, dbFile *model.DBFile, fileBody []byte) error {
	if dbFile == nil {
		return errors.New("No DB file provided")
	}
	if len(fileBody) == 0 {
		return errors.New("No file body provided to upload")
	}
	// Check google bucket exists
	bkt := repo.storage.Bucket(repo.bucket)
	_, err := bkt.Attrs(ctx)
	if err == storage.ErrBucketNotExist {
		return fmt.Errorf("Bucket '%s' doesn't exist", repo.bucket)
	}
	if err != nil {
		return errors.Wrap(err, "repo.storage.Bucket.Attrs")
	}

	obj := bkt.Object(dbFile.Filename)
	objWriter := obj.NewWriter(ctx)
	if _, err := objWriter.Write(fileBody); err != nil {
		return errors.Wrap(err, "objWriter.Write(fileBody)")
	}
	if err := objWriter.Close(); err != nil {
		return errors.Wrap(err, "objWriter.Close()")
	}

	return nil
}

// Download file from Google Cloud storage
func (repo *GoogleCloudFileContentRepo) Download(ctx context.Context, dbFile *model.DBFile) ([]byte, error) {
	if dbFile == nil {
		return nil, errors.New("No DB file provided")
	}
	// Check google bucket exists
	bkt := repo.storage.Bucket(repo.bucket)
	_, err := bkt.Attrs(ctx)
	if err == storage.ErrBucketNotExist {
		return nil, fmt.Errorf("Bucket '%s' doesn't exist", repo.bucket)
	}
	if err != nil {
		return nil, errors.Wrap(err, "repo.storage.Bucket.Attrs")
	}

	obj := bkt.Object(dbFile.Filename)
	attrs, err := obj.Attrs(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "obj.Attrs")
	}
	objReader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "obj.NewReader")
	}

	var buf bytes.Buffer
	if _, err := io.Copy(io.Writer(&buf), objReader); err != nil {
		return nil, err
	}

	dbFile.ContentType = attrs.ContentType

	return buf.Bytes(), nil
}
