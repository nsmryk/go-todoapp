package model

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DBConnection() *sql.DB {
	dsn := GetDBConfig()
	var err error
	//TODO: an error happens here.
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("DB Error1: %w", err))
	}
	CreateTable(db)
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("DB Error2: %w", err))
	}
	return sqlDB
}

func GetDBConfig() string {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, hostname, port, dbname) + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}

func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&Task{})
}
