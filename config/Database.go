package config

import (
	"net"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(appConfig AppConfig) {

	var err error
	var dbURL = appConfig.DatabaseURL

	dbConfig, err := url.Parse(dbURL)
	if err != nil {
		panic(err)
	}

	dbUsername := dbConfig.User.Username()
	dbPassword, _ := dbConfig.User.Password()
	dbHost, dbPort, _ := net.SplitHostPort(dbConfig.Host)

	// var dsn = "root:root@tcp(127.0.0.1:3306)/sppd?charset=utf8mb4&parseTime=True&loc=Local"
	var dsn = dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/sppd?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Tidak dapat terhubung ke database" + dsn)

	}

}
