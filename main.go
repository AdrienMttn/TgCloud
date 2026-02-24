package main

import (
	"fmt"
	"mon-api/models"
	"net/http"
	"os"

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

	http.HandleFunc("/supports-formats", format.GetAcceptFormat)
	http.HandleFunc("/post", tgCloud.PostImage)
	http.HandleFunc("/img/{id}", tgCloud.GetPhotoOnTg)
	
	fmt.Println("Serveur web démarré sur http://localhost:8080")
	http.ListenAndServe(":8100",nil)
	fmt.Println("\nArrêt en cours...")
}

