package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = ""
	dbname   = "stock_db"
)

func main() {
	//สร้างตัวแปรเก็บ database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local")

	//สร้าง log เอาไว้ดู
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  true,          // Disable color
		},
	)

	// สร้าง database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	//ตรวจสอบว่า connect database ได้ไหม
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println(db)
	fmt.Println("Database connection success...")

	// ทำ Migration ตัว model กับ database
	// db.AutoMigrate(&Product{})
	// fmt.Println("Migration success")
}
