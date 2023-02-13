package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(appConfig AppConfig) {

	var err error

	// var dsn = "root:root@tcp(127.0.0.1:3306)/sppd?charset=utf8mb4&parseTime=True&loc=Local"
	// var dsn = appConfig.DatabaseUsername + ":" + appConfig.DatabasePassword + "@tcp(" + appConfig.DatabaseHost + ":" + appConfig.DatabasePort + ")/" + appConfig.DatabaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	var dsn = appConfig.DatabaseURL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Tidak dapat terhubung ke database")
	}

}
