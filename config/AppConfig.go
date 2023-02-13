package config

type AppConfig struct {
	DatabaseUsername string
	DatabasePassword string
	DatabasePort     string
	DatabaseName     string
	AppPort          string
	AllowOrigins     string
}
