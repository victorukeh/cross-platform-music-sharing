package models

import "google.golang.org/api/youtube/v3"

type YouTubePlaylistResponse struct {
	Message   string              `json:"message"`
	Playlists []*youtube.Playlist `json:"playlists"`
}
