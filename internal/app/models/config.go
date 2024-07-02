package models

type AppConfig struct {
	Env  string
	Port string

	LogLevel string

	SpotifyProvider *SpotifyProviderAppConfig
}

type SpotifyProviderAppConfig struct {
	Key         string
	Secret      string
	RedirectURL string
}
