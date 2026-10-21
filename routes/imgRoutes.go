package routes

import (
	"net/http"
	"telegram-cloud/controllers"
)

func registerImg(mux *http.ServeMux, c controllers.ImgController) {
	mux.HandleFunc("POST /img/", c.Post)
	mux.HandleFunc("GET /img/{id}", c.Get)
}