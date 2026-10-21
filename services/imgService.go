package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"telegram-cloud/models"
	"time"

	"github.com/go-telegram/bot"

	tgModels "github.com/go-telegram/bot/models"
)

type ImgService struct {
	Bot *bot.Bot
	ChatId string
}


func (s *ImgService) PostWithUrl(url string, ctx context.Context) (models.Image, *models.ErrorResponse) {
	msg, err := s.Bot.SendDocument(ctx,&bot.SendDocumentParams{
			ChatID: s.ChatId,
			Document: &tgModels.InputFileString{Data: url},
	})
	if err != nil {
		fmt.Print(err)
		return models.Image{ID: ""}, &models.ErrorResponse{ERROR: "Error when uploading on telegram"}
	}
	return models.Image{ID: msg.Document.FileID}, nil
}

func (s *ImgService) PostWithFile(img []byte, filename string, ctx context.Context) (models.Image, *models.ErrorResponse) {
	msg, err := s.Bot.SendDocument(ctx,&bot.SendDocumentParams{
			ChatID: s.ChatId,
			Document: &tgModels.InputFileUpload{Filename: filename, Data: bytes.NewReader(img)},
	})
	if err != nil {
		fmt.Print(err)
		return models.Image{ID: ""}, &models.ErrorResponse{ERROR: ""}
	}
	return models.Image{ID: msg.Document.FileID}, nil
}

func (s *ImgService) Get(id string, ctx context.Context) ([]byte, *models.ErrorResponse) {
	file, err := s.Bot.GetFile(ctx, &bot.GetFileParams{FileID: id})
	if err != nil {
		return []byte{}, &models.ErrorResponse{ERROR: "Error when requesting telegram"}
	}
	url := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", s.Bot.Token(), file.FilePath)
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return []byte{}, &models.ErrorResponse{ERROR: "Error while fetching the image from Telegram"}
	}
	defer resp.Body.Close()
	data , err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, &models.ErrorResponse{ERROR: "Error while reading the image data"}
	}
	return data, nil
}