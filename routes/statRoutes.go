package routes

import (
	"net/http"
	"telegram-cloud/controllers"
)

func registerStat(mux *http.ServeMux, c controllers.StatController) {
	mux.HandleFunc("GET /", c.Get)
}