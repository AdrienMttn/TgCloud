package routes

import (
	"encoding/json"
	"net/http"
	"telegram-cloud/controllers"
)



type Controllers struct {
	Img *controllers.ImgController
	AcceptFormat *controllers.AcceptFormatController
}

func SetupRouter(c Controllers, v string) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to TgCloud API", "version": v})
	})
	registerImg(router, *c.Img)
	registerAcceptFormat(router, *c.AcceptFormat)

	return router
}