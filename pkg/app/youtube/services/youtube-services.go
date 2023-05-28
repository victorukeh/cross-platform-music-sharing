package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

func GetClient(ctx context.Context, config *oauth2.Config) *http.Client {
	tokenFile := "token.json"
	token, err := TokenFromFile(tokenFile)
	if err != nil {
		token = GetTokenFromWeb(config)
		SaveToken(tokenFile, token)
	}
	return config.Client(ctx, token)
}

func GetTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser, then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		panic(err)
	}

	token, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		panic(err)
	}
	return token
}

func TokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	return t, err
}

func SaveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func GetClientSecret() []byte {
	// Replace YOUR_CLIENT_SECRET_FILE with the path to your client_secret.json file.
	clientSecretFile := "client_secret.json"
	b, err := ioutil.ReadFile(clientSecretFile)
	if err != nil {
		panic(err)
	}
	return b
}
