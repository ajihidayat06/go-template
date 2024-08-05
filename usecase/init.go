package usecase

import (
	"go-template/config"
	"go-template/repo"
	usecase "go-template/usecase/book"

	"gorm.io/gorm"
)

type InitUseCaseStruct struct {
	BookUseCase usecase.BookUseCase
}

func InitUseCase(setupRepo *repo.InitRepoStruct, db *gorm.DB, cfg *config.CfgStruct) InitUseCaseStruct {
	return InitUseCaseStruct{
		BookUseCase: usecase.NewBookUseCase(setupRepo.RepoBook, db, cfg),
	}
}
