package config

type AppConfig struct {
	DatabaseHost string
	// DatabaseURL string
	DatabasePassword string
	DatabasePort     string
	DatabaseName     string
	AppPort          string
	AllowOrigins     string
	SecretKey        string
}
