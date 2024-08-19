package config

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	Validator *validator.Validate
	Loging    *LoggerStruct
)

type CfgStruct struct {
}

func NewCfg() CfgStruct {
	return CfgStruct{}
}
