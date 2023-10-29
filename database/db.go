package database

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Database 데이터베이스 구조체
type Database struct {
	*gorm.DB
}

// NewDatabase 생성자
func NewDatabase(dialector gorm.Dialector) *Database {
	config := gorm.Config{TranslateError: true}
	db, err := gorm.Open(dialector, &config)

	if err != nil {
		log.Fatal().Err(err)
	}

	return &Database{db}
}
