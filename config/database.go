package config

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//os.GETENV() main

//os.GETENV("DBHOST")  .env DBHOST=localhost

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "QwE1234$$"
	dbname   = "my_database"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	return db
}
