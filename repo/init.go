package repo

import (
	"go-template/config"

	"gorm.io/gorm"
)

type InitRepoStruct struct {
	cfg      *config.CfgStruct
	RepoBook RepoBook
}

func InitRepo(db *gorm.DB, cfg *config.CfgStruct) InitRepoStruct {
	return InitRepoStruct{
		cfg:      cfg,
		RepoBook: NewRepoBook(db),
	}
}
