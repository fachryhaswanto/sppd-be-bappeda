package config

type AppConfig struct {
	DatabaseHost     string
	DatabaseUsername string
	DatabasePassword string
	DatabasePort     string
	DatabaseName     string
	AppPort          string
	AllowOrigins     string
}
