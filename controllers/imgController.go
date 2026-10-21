package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"telegram-cloud/models"
	"time"
)

type ImgService interface {
	PostWithUrl(url string, ctx context.Context) (models.Image, *models.ErrorResponse)
	PostWithFile(img []byte, filename string, ctx context.Context) (models.Image, *models.ErrorResponse)
	Get(id string, ctx context.Context) ([]byte, *models.ErrorResponse)
}

type ImgController struct {
	service ImgService
}

func NewImgController(s ImgService) *ImgController {
	return &ImgController{service: s}
}

func (c *ImgController) Post(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	img, header, err := r.FormFile("img")
	if (err != nil) {
		data, _ := io.ReadAll(r.Body)
		var body models.Upload;
		err := json.Unmarshal(data, &body);
		if err != nil{
			writeJSON(w, 400, models.ErrorResponse{ERROR: "Invalide request"})
			return
		}
		res, errorResponse :=  c.service.PostWithUrl(body.URL,ctx)
		if errorResponse != nil {
			writeJSON(w, 400, errorResponse)
			return
		}
		writeJSON(w, 200, res)
	} else {
			fileHeader, _ := io.ReadAll(img)
			defer img.Close()
			format := http.DetectContentType(fileHeader)
			AcceptFormat := models.NewAcceptFormat()
			if !AcceptFormat.TestIsSupport(format) {
				writeJSON(w, 400, models.ErrorResponse{ERROR: fmt.Sprintf("This format '%s' is not supported", format)})
				return
			}
			res, errorResponse := c.service.PostWithFile(fileHeader,header.Filename, ctx)
			if errorResponse != nil {
				writeJSON(w, 400, errorResponse)
				return
			}
			writeJSON(w, 200, res)
	}
}

func (c *ImgController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	id := r.PathValue("id")
	if id == "" {
		writeJSON(w, 400, models.ErrorResponse{ERROR: "Empty id"})
		return
	}
	data, err := c.service.Get(id, ctx)
	if err != nil {
		writeJSON(w, 400, err)
	}
	w.Header().Set("Content-Type", http.DetectContentType(data))
	io.Copy(w, bytes.NewReader(data))
}