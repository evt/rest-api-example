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
	tableName   struct{}  `pg:"files"`
	ID          uuid.UUID `pg:"id,notnull,pk"`
	Filename    string    `pg:"filename,notnull"`
	ContentType string    `pg:"-"`
	Created     time.Time `pg:"created,notnull"`
}

// ToWeb converts DBFile to File
func (dbFile *DBFile) ToWeb() *File {
	return &File{
		ID:       dbFile.ID,
		Filename: dbFile.Filename,
		Created:  dbFile.Created,
	}
}
