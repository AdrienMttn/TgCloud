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
	http.HandleFunc("/post-from-file", tgCloud.PostImageFromFile)
	http.HandleFunc("/post-from-url", tgCloud.PostImageFromUrl)
	http.HandleFunc("/img/{id}", tgCloud.GetPhotoOnTg)
	
	fmt.Println("Serveur web démarré sur http://localhost:" + os.Getenv("PORT"))
	http.ListenAndServe(":" + os.Getenv("PORT"),nil)
	fmt.Println("\nArrêt en cours...")
}

