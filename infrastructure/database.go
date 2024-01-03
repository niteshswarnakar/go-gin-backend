package infrastructure

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DBConn struct {
	Db *gorm.DB
}

func NewDbConnection(env Env) DBConn {
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName
	dbHost := env.DBHost
	dbPort := env.DBPort
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Cannot connect to the database", err)
	}
	log.Println("Successfully connected to database", dbName)
	return DBConn{
		Db: db,
	}
}
