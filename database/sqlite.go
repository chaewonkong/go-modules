package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SQLiteDialector(name string) gorm.Dialector {
	dsn := fmt.Sprintf("%s.db", name)

	return sqlite.Open(dsn)
}
