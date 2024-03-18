package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config Config) *gorm.DB {
	user := config.Get("DB_USER")
	password := config.Get("DB_PASSWORD")
	database := config.Get("DB_NAME")
	host := config.Get("DB_HOST")
	port := config.Get("DB_PORT")
	sslMode := config.Get("DB_SSL_MODE")
	dbTypeZone := config.Get("DB_TYPE_ZONE")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=" + sslMode + " TimeZone=" + dbTypeZone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db

}
