package model

import (
	"github.com/google/uuid"
	"time"
)

// User is a JSON user
type User struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Created   time.Time `json:"created"`
}

// ToDB converts User to DBUser
func (user *User) ToDB() *DBUser {
	return &DBUser{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Created:   user.Created,
	}
}

// DBUser is a Postgres user
type DBUser struct {
	tableName struct{}  `pg:"users"`
	ID        uuid.UUID `pg:"id,notnull,pk"`
	Firstname string    `pg:"firstname,notnull"`
	Lastname  string    `pg:"lastname,notnull"`
	Created   time.Time `pg:"created,notnull"`
}

// ToWeb converts DBUser to User
func (dbUser *DBUser) ToWeb() *User {
	return &User{
		ID:        dbUser.ID,
		Firstname: dbUser.Firstname,
		Lastname:  dbUser.Lastname,
		Created:   dbUser.Created,
	}
}
