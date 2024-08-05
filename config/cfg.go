package config

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type CfgStruct struct {
	Validatior *validator.Validate
}

func NewCfg() CfgStruct {
	return CfgStruct{
		Validatior: validator.New(),
	}
}
