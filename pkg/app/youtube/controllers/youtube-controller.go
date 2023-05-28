package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/victorukeh/cross-platform-music-sharing/pkg/app/youtube/config"
	"google.golang.org/api/youtube/v3"
)

func MigrateAllPlaylistsToSpotify(w http.ResponseWriter, r *http.Request) {
	service, err := config.YoutubeConnect()
	if err != nil {
		fmt.Println(err)
	}
	playlists := []*youtube.Playlist{}
	nextPageToken := ""
	for {
		playlistCall := service.Playlists.List([]string{"snippet"}).Mine(true).MaxResults(50).PageToken(nextPageToken)
		playlistResponse, err := playlistCall.Do()
		if err != nil {
			log.Fatalf("Unable to retrieve playlists: %v", err)
		}
		playlists = append(playlists, playlistResponse.Items...)
		nextPageToken = playlistResponse.NextPageToken
		if nextPageToken == "" {
			break
		}
	}
	fmt.Println(playlists)
	// results := &models.YouTubePlaylistResponse{Message: "", Playlists: playlists}
	// res, _ := json.Marshal(playlists)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
	// Search for a song with the given query.
	// query := "never gonna give you up"
	// searchCall := service.Search.List([]string{"id"}).Type("video").Q(query).MaxResults(1)

	// searchResponse, err := searchCall.Do()
	// if err != nil {
	// 	panic(err)
	// }

	// // Get the ID of the first music track in the search results.
	// videoId := searchResponse.Items[0].Id.VideoId

	// // Create a new playlist.
	// playlistTitle := "My Favorite Songs"
	// playlistDescription := "A playlist of my favorite songs"
	// playlistCall := service.Playlists.Insert([]string{"snippet"}, &youtube.Playlist{
	// 	Snippet: &youtube.PlaylistSnippet{
	// 		Title:       playlistTitle,
	// 		Description: playlistDescription,
	// 	},
	// })

	// playlistResponse, err := playlistCall.Do()
	// if err != nil {
	// 	panic(err)
	// }

	// // Add the music track to the new playlist.
	// playlistItemId := youtube.PlaylistItemSnippet{
	// 	PlaylistId: playlistResponse.Id,
	// 	ResourceId: &youtube.ResourceId{
	// 		Kind:    "youtube#video",
	// 		VideoId: videoId,
	// 	},
	// }
	// playlistItemCall := service.PlaylistItems.Insert([]string{"snippet"}, &youtube.PlaylistItem{
	// 	Snippet: &playlistItemId,
	// })

	// _, err = playlistItemCall.Do()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Successfully added song to playlist.\n")
}
