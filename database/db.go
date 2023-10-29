package database

import (
	"gorm.io/gorm"
)

// NewDatabase 생성자
func NewDatabase(dialector gorm.Dialector) (db *gorm.DB, err error) {
	config := gorm.Config{TranslateError: true}
	db, err = gorm.Open(dialector, &config)

	if err != nil {
		return
	}

	return
}
