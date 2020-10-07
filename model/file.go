package model

import (
	"time"

	"github.com/google/uuid"
)

// File holds file metadata as a JSON
type File struct {
	ID       uuid.UUID `json:"id"`
	Filename string    `json:"filename" validate:"required"`
	Created  time.Time `json:"created"`
}

// ToDB converts File to DBFile
func (file *File) ToDB() *DBFile {
	return &DBFile{
		ID:       file.ID,
		Filename: file.Filename,
		Created:  file.Created,
	}
}

// DBFile is a file in Postgres
type DBFile struct {
	tableName   struct{}  `pgdb:"files"`
	ID          uuid.UUID `pgdb:"id,notnull,pk"`
	Filename    string    `pgdb:"filename,notnull"`
	ContentType string    `pgdb:"-"`
	Created     time.Time `pgdb:"created,notnull"`
}

// ToWeb converts DBFile to File
func (dbFile *DBFile) ToWeb() *File {
	return &File{
		ID:       dbFile.ID,
		Filename: dbFile.Filename,
		Created:  dbFile.Created,
	}
}
