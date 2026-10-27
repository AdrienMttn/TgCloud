package services

import (
	"telegram-cloud/models"
	"time"
)

type StatService struct {
	v string
	start_time time.Time
}

func NewStatService(version string) *StatService {
	return &StatService{v: version, start_time: time.Now()}
}

func (s *StatService) Get() (models.Stat, *models.ErrorResponse) {
	stat := models.Stat{
		Version: s.v,
		Uptime_Seconds: int(time.Since(s.start_time).Seconds()),
		Started_At: s.start_time.Format(time.ANSIC),
	}
	return stat, nil
}

func (s *StatService) ImageCount() int {
	return 0
}