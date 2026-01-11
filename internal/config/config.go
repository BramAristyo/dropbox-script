package config

import "os"

type Config struct {
	Local   LocalConfig
	Dropbox DropboxConfig
}

type LocalConfig struct {
	Path string
}

type DropboxConfig struct {
	AppKey       string
	SecretKey    string
	RefreshToken string
	Path         string
}

func GetConfig() *Config {
	cfg := &Config{}
	cfg.Dropbox = DropboxConfig{
		AppKey:       os.Getenv("DROPBOX_APP_KEY"),
		SecretKey:    os.Getenv("DROPBOX_APP_SECRET"),
		RefreshToken: os.Getenv("DROPBOX_REFRESH_TOKEN"),
		Path:         os.Getenv("DROPBOX_PATH"),
	}

	cfg.Local = LocalConfig{
		Path: os.Getenv("LOCAL_PATH"),
	}

	return cfg
}
