package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Telegram TelegramConfig
	Api ApiConfig
}

type ApiConfig struct {
	Port string
}
type TelegramConfig struct {
	BotToken string
	ChatId   string
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")
	
	config := &Config{
		Telegram: TelegramConfig{
			BotToken: os.Getenv("BOT_TOKEN"),
			ChatId: os.Getenv("CHAT_ID"),
		},
		Api: ApiConfig{
			Port: os.Getenv("PORT"),
		},
	}

	return config
}