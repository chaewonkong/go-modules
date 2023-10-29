package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySQLDialector(user string, password string, url string, port string, name string) gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		url,
		port,
		name,
	)
	return mysql.Open(dsn)
}
