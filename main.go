package main

import (
	"fmt"
	"mon-api/models"
	"net/http"
	"os"

	"github.com/go-telegram/bot"
)

var AcceptTypeImg = [4]string{"image/jpeg","image/png","image/jpg","image/gif"}

type StatusApi struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

type AcceptFormat struct {
	Supports [4]string `json:"supports"`
}


func main() {

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
	http.ListenAndServe(":8080",nil)
	fmt.Println("\nArrêt en cours...")
}

