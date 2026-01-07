package config

import "os"

type Config struct {
	Dropbox DropboxConfig
}

type DropboxConfig struct {
	AppKey       string
	SecretKey    string
	RefreshToken string
}

func GetConfig() *Config {
	cfg := &Config{}
	cfg.Dropbox = DropboxConfig{
		AppKey:       os.Getenv("DROPBOX_APP_KEY"),
		SecretKey:    os.Getenv("DROPBOX_APP_SECRET"),
		RefreshToken: os.Getenv("DROPBOX_REFRESH_TOKEN"),
	}

	return cfg
}
