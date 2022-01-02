package gateways

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"golang.org/x/oauth2"
)

const (
	_appNubankBucket     = "APP_NUBANK_BUCKET"
	_appNubankSecretFile = "APP_NUBANK_SECRET_FILE"
)

type googleToken struct {
	RefreshToken string `json:"refresh_token"`
	Token        string `json:"token"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	TokenURI     string `json:"token_uri"`
}

func NewGoogleToken() *googleToken {
	content, err := ReadContentFile(os.Getenv(_appNubankBucket), os.Getenv(_appNubankSecretFile))

	g := googleToken{}

	if err != nil {
		return &g
	}

	json.Unmarshal([]byte(content), &g)
	return &g
}

func (g *googleToken) GetTokenSource() oauth2.TokenSource {
	ctx := context.Background()

	token := &oauth2.Token{
		RefreshToken: g.RefreshToken,
		AccessToken:  g.Token,
		Expiry:       time.Now().Add(time.Duration(-10) * time.Minute),
		TokenType:    "Bearer",
	}

	config := &oauth2.Config{
		ClientID:     g.ClientID,
		ClientSecret: g.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: g.TokenURI,
		},
	}

	return config.TokenSource(ctx, token)
}
