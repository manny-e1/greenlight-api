package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.Handler(http.MethodGet, "/v1/healthcheck", http.HandlerFunc(app.healthCheckHandler))
	router.Handler(http.MethodPost, "/v1/movies", http.HandlerFunc(app.createMovieHandler))
	router.Handler(http.MethodGet, "/v1/movies/:id", http.HandlerFunc(app.showMovieHandler))

	return router
}
