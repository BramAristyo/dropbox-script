package main

import (
	"fmt"
	"time"

	"github.com/BramAristyo/dropbox-script/internal/config"
	"github.com/BramAristyo/dropbox-script/internal/dropbox"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.GetConfig()

	token, err := dropbox.GetNewToken(cfg.Dropbox.AppKey, cfg.Dropbox.SecretKey, cfg.Dropbox.RefreshToken)

	if err != nil {
		fmt.Println("Failed get Dropbox token: ", err)
	}

	dropbox.Sync(token.AccessToken, cfg.Dropbox.Path, cfg.Local.Path)

	loc, err := time.LoadLocation("Asia/Jakarta")
	fmt.Printf(" %s\n", time.Now().In(loc))
}
