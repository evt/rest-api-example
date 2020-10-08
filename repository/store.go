package repository

// Store contains all repositories
type Store struct {
	User        UserRepo
	File        FileRepo
	FileContent FileContentRepo
}
