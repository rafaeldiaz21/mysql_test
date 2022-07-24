package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var globalDB *gorm.DB

func NewConnection() *gorm.DB {
	var err error
	globalDB, err = dbConnection()

	if err != nil {
		fmt.Print(err)
		panic("Failed to connect database")
	}

	return globalDB
}

// Se debe configurar para leer las variables de entorno
func dbConnection() (*gorm.DB, error) {
	var err error
	var dsn string
	if globalDB != nil {
		return globalDB, err
	}

	// docker connection
	dsn = "tester:secret@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	// local connection
	// dsn = "root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	globalDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return globalDB, nil
}
