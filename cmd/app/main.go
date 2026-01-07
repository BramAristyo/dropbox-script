package main

import (
	"fmt"

	"github.com/BramAristyo/dropbox-script/internal/config"
	"github.com/BramAristyo/dropbox-script/internal/dropbox"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.GetConfig()

	token, err := dropbox.GetNewToken(cfg.Dropbox.AppKey, cfg.Dropbox.SecretKey, cfg.Dropbox.RefreshToken)

	if err != err {
		fmt.Println("Failed get Dropbox token: ", err)
	}

	fmt.Println(token.AccessToken)
}
