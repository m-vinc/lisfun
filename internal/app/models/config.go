package models

type AppConfig struct {
	SecretKey string

	Env  string
	Port string

	LogLevel string

	SpotifyProvider *SpotifyProviderAppConfig

	DatabaseURL string
}

type SpotifyProviderAppConfig struct {
	Key         string
	Secret      string
	RedirectURL string
}
