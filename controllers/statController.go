package controllers

import (
	"net/http"
	"telegram-cloud/models"
)

type StatService interface {
	Get() (models.Stat, *models.ErrorResponse)
}

type StatController struct {
	service StatService
}

func NewStatController(s StatService) *StatController{
	return &StatController{service: s}
}

func (c *StatController) Get(w http.ResponseWriter, r *http.Request) {
	res, err := c.service.Get()
	if (err != nil) {
		writeJSON(w, 400, err)
		return
	}
	writeJSON(w, 200, res)
}