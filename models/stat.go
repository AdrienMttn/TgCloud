package models


type Stat struct {
	Version string `json:"version"`
	Uptime_Seconds int `json:"uptime_seconds"`
	Started_At string `json:"started_at"`
}