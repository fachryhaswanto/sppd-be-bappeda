package config

type AppConfig struct {
	DatabaseUsername string
	DatabasePassword string
	DatabasePort     string
	DatabaseHost     string
	DatabaseName     string
	AppPort          string
	SecretKey        string
}
