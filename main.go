package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"telegram-cloud/models"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load(".env")

	b, err := bot.New(os.Getenv("BOT_TOKEN"))
	tgCloud := models.TgCloud{Bot:b, ChatId: os.Getenv("CHAT_ID")}
	format := models.NewAcceptFormat()

	if nil != err {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to TgCloud API", "version": "2.1"})
	})
	http.HandleFunc("/supports-formats", format.GetAcceptFormat)
	http.HandleFunc("/post-from-file", tgCloud.PostImageFromFile)
	http.HandleFunc("/post-from-url", tgCloud.PostImageFromUrl)
	http.HandleFunc("/img/{id}", tgCloud.GetPhotoOnTg)
	
	fmt.Println("Serveur web démarré sur http://localhost:" + os.Getenv("PORT"))
	http.ListenAndServe(":" + os.Getenv("PORT"),nil)
	fmt.Println("\nArrêt en cours...")
}

