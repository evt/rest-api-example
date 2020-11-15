package local

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/evt/rest-api-example/model"
	"github.com/pkg/errors"
)

// GoogleCloudFileContentRepo ...
type GoogleCloudFileContentRepo struct {
	filePath string
}

// NewFileContentRepo ...
func NewFileContentRepo(filePath string) *GoogleCloudFileContentRepo {
	return &GoogleCloudFileContentRepo{filePath: filePath}
}

// Upload file to Google Cloud storage
func (repo *GoogleCloudFileContentRepo) Upload(ctx context.Context, dbFile *model.DBFile, fileBody []byte) error {
	if dbFile == nil {
		return errors.New("No DB file provided")
	}

	if len(fileBody) == 0 {
		return errors.New("No file body provided to upload")
	}

	if err := os.MkdirAll(repo.filePath, os.ModePerm); err != nil {
		return errors.Wrap(err, "os.MkdirAll failed")
	}

	return ioutil.WriteFile(repo.filePath+"/"+dbFile.Filename, fileBody, 0644)
}

// Download file from Google Cloud storage
func (repo *GoogleCloudFileContentRepo) Download(ctx context.Context, dbFile *model.DBFile) ([]byte, error) {
	if dbFile == nil {
		return nil, errors.New("No DB file provided")
	}

	return ioutil.ReadFile(repo.filePath + "/" + dbFile.Filename)
}
