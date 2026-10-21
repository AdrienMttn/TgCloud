package routes

import (
	"net/http"
	"telegram-cloud/controllers"
)

func registerAcceptFormat(mux *http.ServeMux, c controllers.AcceptFormatController) {
	mux.HandleFunc("GET /format/", c.Get)
}