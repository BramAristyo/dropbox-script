package dropbox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/BramAristyo/dropbox-script/models"
)

func GetNewToken(appKey, appSecret, refreshToken string) (models.Token, error) {
	v := url.Values{}
	v.Set("grant_type", "refresh_token")
	v.Set("refresh_token", refreshToken)
	v.Set("client_id", appKey)
	v.Set("client_secret", appSecret)

	url := "https://api.dropbox.com/oauth2/token"

	hc := http.Client{}
	r, err := http.NewRequest("POST", url, strings.NewReader(v.Encode()))

	if err != nil {
		fmt.Println("Error: failed make Request")
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := hc.Do(r)

	token := models.Token{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&token)

	return token, err
}
