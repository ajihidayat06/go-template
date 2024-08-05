package repo

import (
	"go-template/model"

	"gorm.io/gorm"
)

type RepoBook interface {
	Insert(tx *gorm.DB, book model.Book) error
}

type RepoBookImpl struct {
	DB *gorm.DB
}

func NewRepoBook(database *gorm.DB) RepoBook {
	return &RepoBookImpl{
		DB: database,
	}
}

func (r *RepoBookImpl) Insert(tx *gorm.DB, book model.Book) error {
	result := tx.Create(&book)
	return result.Error
}
