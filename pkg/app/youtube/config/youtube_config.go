package config

import (
	"context"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func YoutubeConnect() (*youtube.Service, error) {
	apiKey := os.Getenv("API_KEY")
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
	}
	// ctx := context.Background()

	// // Replace YOUR_API_KEY with your actual API key.
	// // Replace YOUR_CLIENT_SECRET_FILE with the path to your client_secret.json file.
	// config, err := google.ConfigFromJSON(services.GetClientSecret(), youtube.YoutubeForceSslScope)
	// if err != nil {
	// 	panic(err)
	// }

	// // client := services.GetClient(ctx, config)

	// // Create a resource for the YouTube Music API.
	// service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	// if err != nil {
	// 	panic(err)
	// }
	return service, err
}
