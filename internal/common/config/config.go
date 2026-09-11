package config

import "os"

type Config struct {
	ServerPort string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	JwtSecret  string
}

func Load() *Config {

	return &Config{

		ServerPort: os.Getenv("SERVER_PORT"),
		DbHost:     os.Getenv("POSTGRES_HOST"),
		DbPort:     os.Getenv("POSTGRES_PORT"),
		DbName:     os.Getenv("POSTGRES_DB"),
		DbUser:     os.Getenv("POSTGRES_USER"),
		DbPassword: os.Getenv("POSTGRES_PASSWORD"),
		JwtSecret:  os.Getenv("JWT_SECRET"),
	}

}
