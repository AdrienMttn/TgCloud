package models

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	. "telegram-cloud/types"
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
	w.Header().Set("Content-Type", "application/json")

	// Log the request details
	fmt.Printf("Action: Image Upload From File | Method: %s\n", r.Method)
	fmt.Printf("USER IP: %s | USER AGENT: %s\n", r.RemoteAddr, r.UserAgent())
	// END LOG
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400,Message: "Method not allowed"})
		// Log the response details
		fmt.Printf("Response: Method not allowed: %s\n", r.Method)
		// END LOG
		return
	}
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400,Message: "Error while parsing the form data"})
		// Log the response details
		fmt.Printf("Response: Error while parsing the form data: %s\n", err.Error())
		// END LOG
		return
	}
	img, header, err := r.FormFile("img")
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400,Message: "Image file is required"})
		// Log the response details
		fmt.Printf("Response: Image file is required\n")
		// END LOG
		return
	}
	fileHeader, _ := io.ReadAll(img)
	defer img.Close()
	var format string 
	format = http.DetectContentType(fileHeader)
	AcceptFormat := NewAcceptFormat()
	
	if !AcceptFormat.TestIsSupport(format) {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400,Message: fmt.Sprintf("This format '%s' is not supported", format)})
		// Log the response details
		fmt.Printf("Response: This format '%s' is not supported\n", format)
		// END LOG
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	response := tgCloud.PostOnTgWithFile(fileHeader, header.Filename, ctx)
	if response.Status != 200 {
		// Log the response details
		fmt.Printf("Response: Error while uploading the image on Telegram: %s\n", response.Message)
		// END LOG
		w.WriteHeader(response.Status)
	}
	// Log the response details
	fmt.Printf("Response: %s\n", response.Message)
	// END LOG
	json.NewEncoder(w).Encode(response)
}

func (tgCloud *TgCloud) PostImageFromUrl(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	// Log the request details
	fmt.Printf("Action: Image Upload From URL | Method: %s\n", r.Method)
	fmt.Printf("USER IP: %s | USER AGENT: %s\n", r.RemoteAddr, r.UserAgent())
	// END LOG
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		// Log the response details
		fmt.Printf("Response: Preflight OPTIONS request handled successfully\n")
		// END LOG
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		// Log the response details
		fmt.Printf("Response: Method not allowed: %s\n", r.Method)
		// END LOG
		json.NewEncoder(w).Encode(Response{Status: 400,Message: "Method not allowed"})
		return
	}
	jsonBody ,_ := io.ReadAll(r.Body)
	var data jsonData
	err := json.Unmarshal(jsonBody, &data)
	if err != nil{
		w.WriteHeader(400)
		// Log the response details
		fmt.Printf("Response: Invalid JSON body: %s\n", err.Error())
		// END LOG
		json.NewEncoder(w).Encode(Response{Status: 400,Message: "Invalid JSON body"})
		return
	}
	response := tgCloud.PostOnTgWithUrl(data.Url, r.Context())
	if response.Status != 200 {
		w.WriteHeader(response.Status)
		// Log the response details
		fmt.Printf("Response: Error while uploading the image on Telegram: %s\n", response.Message)
		// END LOG
	}
	// Log the response details
	fmt.Printf("Response: %s\n", response.Message)
	// END LOG
	json.NewEncoder(w).Encode(response)
}


func (tgCloud *TgCloud) PostOnTgWithFile(img []byte, fileName string, ctx context.Context) Response {
	msg, err := tgCloud.Bot.SendDocument(ctx,&bot.SendDocumentParams{
			ChatID: tgCloud.ChatId,
			Document: &models.InputFileUpload{Filename: fileName, Data: bytes.NewReader(img)},
	})
	if err != nil {
		fmt.Print(err)
		return Response{Status: 400, Message: err.Error()}
	}
	return Response{Status: 200, Message: "Image uploaded successfully. Keep this ID safe.", Id:msg.Document.FileID}
}

func (tgCloud *TgCloud) PostOnTgWithUrl(url string, ctx context.Context) Response {
	msg, err := tgCloud.Bot.SendDocument(ctx,&bot.SendDocumentParams{
			ChatID: tgCloud.ChatId,
			Document: &models.InputFileString{Data: url},
	})
	if err != nil {
		fmt.Print(err)
		return Response{Status: 400, Message: err.Error()}
	}
	return Response{Status: 200, Message: "Image uploaded successfully. Keep this ID safe.", Id: msg.Document.FileID}
}


func (tgCloud *TgCloud) GetPhotoOnTg(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Log the request details
	fmt.Printf("Action: Get Photo On Telegram | Method: %s\n", r.Method)
	fmt.Printf("USER IP: %s | USER AGENT: %s\n", r.RemoteAddr, r.UserAgent())
	// END LOG
	if r.Method != "GET"{
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400, Message:"Method not allowed"})
		// Log the response details
		fmt.Printf("Response: Method not allowed: %s\n", r.Method)
		// END LOG
		return
	}
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400, Message:"Empty id"})
		// Log the response details
		fmt.Printf("Response: Empty id\n")
		// END LOG
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	file, err := tgCloud.Bot.GetFile(ctx, &bot.GetFileParams{FileID: id})
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(Response{Status: 400, Message:err.Error()})
		// Log the response details
		fmt.Printf("Response: Error while fetching the file from Telegram: %s\n", err.Error())
		// END LOG
		return
	}
	tgCloud.Bot.GetFile(ctx, &bot.GetFileParams{FileID: id})
	url := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", os.Getenv("BOT_TOKEN"), file.FilePath)
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		json.NewEncoder(w).Encode(Response{Status: 400,Message: "Error while fetching the image from Telegram"})
		// Log the response details
		fmt.Printf("Response: Error while fetching the image from Telegram: %s\n", err.Error())
		// END LOG
		return
	}
	defer resp.Body.Close()
	data , err := io.ReadAll(resp.Body)
	if err != nil {
		json.NewEncoder(w).Encode(Response{Status: 400,Message: "Error while reading the image data"})
		// Log the response details
		fmt.Printf("Response: Error while reading the image data: %s\n", err.Error())
		// END LOG
		return
	}
	fmt.Printf("Image fetched successfully from Telegram, size: %d bytes\n", len(data))
	w.Header().Set("Content-Type", http.DetectContentType(data))
	io.Copy(w, bytes.NewReader(data))
}


