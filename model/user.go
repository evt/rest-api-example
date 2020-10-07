package model

import (
	"time"

	"github.com/google/uuid"
)

// User is a JSON user
type User struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname" validate:"required"`
	Lastname  string    `json:"lastname" validate:"required"`
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
	tableName struct{}  `pgdb:"users"`
	ID        uuid.UUID `pgdb:"id,notnull,pk"`
	Firstname string    `pgdb:"firstname,notnull"`
	Lastname  string    `pgdb:"lastname,notnull"`
	Created   time.Time `pgdb:"created,notnull"`
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
