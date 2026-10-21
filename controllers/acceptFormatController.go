package controllers

import (
	"net/http"
	"telegram-cloud/models"
)

type AcceptFormatService interface{
	Get() (models.AcceptFormat)
}

type AcceptFormatController struct {
	service AcceptFormatService
}

func NewAcceptFormatController(s AcceptFormatService) (*AcceptFormatController) {
	return &AcceptFormatController{service: s}
}

func (c *AcceptFormatController) Get(w http.ResponseWriter, r *http.Request) {
	writeJSON(w,200,c.service.Get())
}