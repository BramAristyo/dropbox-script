package main

import (
	"fmt"
	"time"

	"github.com/BramAristyo/dropbox-script/internal/config"
	"github.com/BramAristyo/dropbox-script/internal/dropbox"
	"github.com/joho/godotenv"
)

func main() {
	start := time.Now()

	_ = godotenv.Load()
	cfg := config.GetConfig()

	fmt.Println("Configuration loaded successfully. Script is starting...")

	token, err := dropbox.GetNewToken(cfg.Dropbox.AppKey, cfg.Dropbox.SecretKey, cfg.Dropbox.RefreshToken)

	fmt.Println("Success Auth .. ")

	if err != nil {
		fmt.Println("Failed get Dropbox token: ", err)
	}

	dropbox.Sync(token.AccessToken, cfg.Dropbox.Path, cfg.Local.Path)

	fmt.Printf("Total runtime: %v\n", time.Since(start))
}
