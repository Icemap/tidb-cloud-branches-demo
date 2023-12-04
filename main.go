package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

// Get the env variables from GitHub workflow
func getDSN() string {
	tidbUser := os.Getenv("USERNAME")
	tidbPassword := os.Getenv("PASSWORD")
	tidbHost := os.Getenv("HOST")
	tidbPort := os.Getenv("PORT")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/test?charset=utf8mb4&tls=true",
		tidbUser, tidbPassword, tidbHost, tidbPort)
}

func createDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(getDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	db := createDB()
	var version string
	db.Raw("SELECT VERSION();").Scan(&version)

	fmt.Println("TiDB version: ", version)
}
