package sqldb

import (
	"gorm.io/gorm"
)

func New(db *gorm.DB) *repo {
	return &repo{db}
}

type repo struct {
	db *gorm.DB
}
