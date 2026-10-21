package services

import "telegram-cloud/models"

type AcceptFormatService struct {
}

func (s *AcceptFormatService) Get() models.AcceptFormat {
	return models.NewAcceptFormat()
}