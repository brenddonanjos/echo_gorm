package database

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_NAME = "echo_gorm"
	DB_USER = "user"
	DB_PASS = "user"
)

var DB *gorm.DB

func StartGorm(params ...string) *gorm.DB {
	var err error

	dsn := DB_USER + ":" + DB_PASS + "@tcp(mysql:3306)/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	return db
}

// Start open connection use sql package
func Start() *sql.DB {
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp(127.0.0.1:3306)/"+DB_NAME)
	if err != nil {
		panic(err)
	}
	return db
}
