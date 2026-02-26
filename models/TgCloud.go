package models

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	. "mon-api/types"
	"net/http"
	"os"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)


type TgCloud struct {
	Bot *bot.Bot
	ChatId string
}

type jsonData struct {
	Url string `json:"url" binding:"required"`
}

func (tgCloud *TgCloud) PostImageFromFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400,Message: "Méthode non autorisée"})
		return
	}
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400,Message: "Formulaire invalide"})
		return
	}
	img, header, err := r.FormFile("img")
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400,Message: "Pas de fichier ou image invalide"})
		return
	}
	fileHeader, _ := io.ReadAll(img)
	defer img.Close()
	var format string 
	format = http.DetectContentType(fileHeader)
	AcceptFormat := NewAcceptFormat()
	
	if !AcceptFormat.TestIsSupport(format) {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400,Message: fmt.Sprintf("Le format '%s' n'est pas supporté", format)})
		return
	}

	// response := tgCloud.Post(fileHeader, header.Filename)
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	response := tgCloud.PostOnTgWithFile(fileHeader, header.Filename, ctx)
	if response.Status != 200 {
		w.WriteHeader(response.Status)
	}
	json.NewEncoder(w).Encode(response)
}

func (tgCloud *TgCloud) PostImageFromUrl(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400,Message: "Méthode non autorisée"})
		return
	}
	jsonBody ,_ := io.ReadAll(r.Body)
	var data jsonData
	err := json.Unmarshal(jsonBody, &data)
	if err != nil{
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400,Message: "Corps de la requête invalide"})
		return
	}
	json.NewEncoder(w).Encode(tgCloud.PostOnTgWithUrl(data.Url, r.Context()))
}


func (tgCloud *TgCloud) PostOnTgWithFile(img []byte, fileName string, ctx context.Context) Response {
	msg, err := tgCloud.Bot.SendPhoto(ctx,&bot.SendPhotoParams{
			ChatID: tgCloud.ChatId,
			Photo: &models.InputFileUpload{Filename: fileName, Data: bytes.NewReader(img)},
	})
	if err != nil {
		fmt.Print(err)
		return Response{Status: 400, Message: err.Error()}
	}
	return Response{Status: 200, Message: "Conserver cette Id", Id: msg.Photo[len(msg.Photo)-1].FileID}
}

func (tgCloud *TgCloud) PostOnTgWithUrl(url string, ctx context.Context) Response {
	msg, err := tgCloud.Bot.SendPhoto(ctx,&bot.SendPhotoParams{
			ChatID: tgCloud.ChatId,
			Photo: &models.InputFileString{Data: url},
	})
	if err != nil {
		fmt.Print(err)
		return Response{Status: 400, Message: err.Error()}
	}
	return Response{Status: 200, Message: "Conserver cette Id", Id: msg.Photo[len(msg.Photo)-1].FileID}
}


func (tgCloud *TgCloud) GetPhotoOnTg(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400, Message:"Méthode non autorisée"})
	}
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400, Message:"Entrer un id"})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	file, err := tgCloud.Bot.GetFile(ctx, &bot.GetFileParams{FileID: id})
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400, Message:err.Error()})
		return
	}
	url := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", os.Getenv("BOT_TOKEN"), file.FilePath)
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		json.NewEncoder(w).Encode(Response{Status: 400,Message: "Erreur lors de la récupération de l'image"})
		return
	}
	defer resp.Body.Close()
	data , err := io.ReadAll(resp.Body)
	if err != nil {
		json.NewEncoder(w).Encode(Response{Status: 400,Message: "Erreur lors de la lecture de l'image"})
		return
	}
	w.Header().Set("Content-Type", http.DetectContentType(data))
	w.Write(data)
}


