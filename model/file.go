package model

import (
	"time"

	"github.com/google/uuid"
)

// File holds file metadata as a JSON
type File struct {
	ID        uuid.UUID `json:"id"`
	Filename  string    `json:"filename" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

// ToDB converts File to DBFile
func (file *File) ToDB() *DBFile {
	return &DBFile{
		ID:        file.ID,
		Filename:  file.Filename,
		CreatedAt: file.CreatedAt,
	}
}

// DBFile is a file in Postgres
type DBFile struct {
	tableName   struct{}  `pg:"files"`
	ID          uuid.UUID `pg:"id,notnull,pk"`
	Filename    string    `pg:"filename,notnull"`
	ContentType string    `pg:"-" gorm:"-"`
	CreatedAt   time.Time `pg:"created_at,notnull"`
}

// TableName overrides default table name for gorm
func (DBFile) TableName() string {
	return "files"
}

// ToWeb converts DBFile to File
func (dbFile *DBFile) ToWeb() *File {
	return &File{
		ID:        dbFile.ID,
		Filename:  dbFile.Filename,
		CreatedAt: dbFile.CreatedAt,
	}
}
