package db

import (
	"arno/configs"
	"fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Dushanbe",
		config.DBConfig.Database.Host,
		config.DBConfig.Database.User,
		config.DBConfig.Database.Password,
		config.DBConfig.Database.DBName,
		config.DBConfig.Database.Port,
		config.DBConfig.Database.SSLMode)

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}
	return db, nil
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	err = dbSQL.Close()
	if err != nil {
		return
	}
}
