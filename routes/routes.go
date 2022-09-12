package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	EpisodeRoutes(r)
	UserRoutes(r)
	FilmRoutes(r)
	CategoryRoutes(r)
	TransactionRoutes(r)
}
