package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"telegram-cloud/config"
	"telegram-cloud/controllers"
	"telegram-cloud/routes"
	"telegram-cloud/services"

	"github.com/go-telegram/bot"
)
var version = "dev"
func main() {
	var CONFIG config.Config = *config.LoadConfig()

	b, err := bot.New(CONFIG.Telegram.BotToken)
	if nil != err {
		panic(err)
	}
	CONTROLLERS := routes.Controllers{
		Img: controllers.NewImgController(&services.ImgService{Bot:b, ChatId: CONFIG.Telegram.ChatId}),
		AcceptFormat: controllers.NewAcceptFormatController(&services.AcceptFormatService{}),
		Stat: controllers.NewStatController(services.NewStatService(version)),
	}
	
	fmt.Println("Serveur web démarré sur http://localhost:" + CONFIG.Api.Port)
	http.ListenAndServe(":" + CONFIG.Api.Port, cors(routes.SetupRouter(CONTROLLERS)))
	fmt.Println("\nArrêt en cours...")
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("Access-Control-Allow-Origin", "*")
		h.Set("Access-Control-Allow-Methods", "*")
		h.Set("Access-Control-Allow-Headers", "*")
		h.Set("Access-Control-Expose-Headers", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}