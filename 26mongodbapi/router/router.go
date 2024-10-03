package router

import (
	"net/http"

	"github.com/ABHI2598/mongoapi/controller"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/movies", controller.GetAllMovies)
	mux.HandleFunc("GET /api/movie/{id}", controller.GetOneMovie)
	mux.HandleFunc("POST /api/movie", controller.CreateMovies)
	mux.HandleFunc("PUT /api/movie/{id}", controller.UpdateOneMovie)
	mux.HandleFunc("DELETE /api/movie/{id}", controller.DeleteOneMovie)

	return mux
}
