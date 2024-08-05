package book

import (
	"context"
	"go-template/config"
	"go-template/errutils"
	"go-template/model"
	"go-template/repo"

	"gorm.io/gorm"
)

type BookUseCase interface {
	InsertBook(ctx context.Context, request model.BookRequest) (response interface{}, err errutils.ErrorModel)
}

type bookUseCaseImpl struct {
	cfg            *config.CfgStruct
	BookRepository repo.RepoBook
	DB             *gorm.DB
}

func NewBookUseCase(repoBook repo.RepoBook, db *gorm.DB, cfg *config.CfgStruct) BookUseCase {
	return &bookUseCaseImpl{
		cfg:            cfg,
		BookRepository: repoBook,
		DB:             db,
	}
}
