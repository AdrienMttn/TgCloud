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
	}

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to TgCloud API", "version": version})
	// })
	// http.HandleFunc("/supports-formats", format.GetAcceptFormat)
	// http.HandleFunc("/post-from-file", tgCloud.PostImageFromFile)
	// http.HandleFunc("/post-from-url", tgCloud.PostImageFromUrl)
	// http.HandleFunc("/img/{id}", tgCloud.GetPhotoOnTg)
	
	fmt.Println("Serveur web démarré sur http://localhost:" + CONFIG.Api.Port)
	http.ListenAndServe(":" + CONFIG.Api.Port, routes.SetupRouter(CONTROLLERS, version))
	fmt.Println("\nArrêt en cours...")
}

