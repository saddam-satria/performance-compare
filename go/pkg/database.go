package pkg

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Connection *gorm.DB

type DatabaseConnection struct {
	HOST     string
	USER     string
	PASSWORD string
	DATABASE string
	PORT     string
}

func Connect(ctx *DatabaseConnection) error {
	dsn := "host=" + ctx.HOST + " user=" + ctx.USER + " password=" + ctx.PASSWORD + " dbname=" + ctx.DATABASE + " port=" + ctx.PORT + " sslmode=disable"
	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if error != nil {
		return error
	}

	Connection = db
	return error
}
