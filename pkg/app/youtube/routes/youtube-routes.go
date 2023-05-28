package routes

import (
	"github.com/gorilla/mux"
	"github.com/victorukeh/cross-platform-music-sharing/pkg/app/youtube/controllers"
)

var YouTubeRoutes = func(router *mux.Router) {
	router.HandleFunc("/youtube/playlists", controllers.MigrateAllPlaylistsToSpotify).Methods(("GET"))
}
