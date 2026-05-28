package config

import "os"

type Config struct {
	AppEnv             string
	HTTPAddr           string
	LogLevel           string
	DatabaseURL        string
	JWTSecret          string
	AccessTokenMinutes int
	RefreshTokenHours  int
}

func Load() Config {
	return Config{
		AppEnv:             getEnv("APP_ENV", "development"),
		HTTPAddr:           getEnv("HTTP_ADDR", ":8080"),
		LogLevel:           getEnv("LOG_LEVEL", "info"),
		DatabaseURL:        getEnv("DATABASE_URL", "postgres://creditoreal:change_me_local_password@localhost:5432/creditoreal_crm?sslmode=disable"),
		JWTSecret:          getEnv("JWT_SECRET", "change_me_dev_jwt_secret"),
		AccessTokenMinutes: getEnvInt("ACCESS_TOKEN_MINUTES", 15),
		RefreshTokenHours:  getEnvInt("REFRESH_TOKEN_HOURS", 168),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed := 0
	for _, char := range value {
		if char < '0' || char > '9' {
			return fallback
		}
		parsed = parsed*10 + int(char-'0')
	}

	if parsed == 0 {
		return fallback
	}

	return parsed
}
