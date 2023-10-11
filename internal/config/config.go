package config

import (
	"log"
	"os"
)

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	LOKI_URL           string
	TELEGRAM_CHAT_NAME string
	LOKI_LOGIN         string
	LOKI_PASSWORD      string
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {
	Settings = SettingsINI{}
	Settings.TELEGRAM_CHAT_NAME = os.Getenv("TELEGRAM_CHAT_NAME")
	Settings.LOKI_URL = os.Getenv("LOKI_URL")

	if Settings.TELEGRAM_CHAT_NAME == "" {
		log.Panic("Error: Need fill TELEGRAM_CHAT_NAME")
	}

	if Settings.LOKI_URL == "" {
		log.Panic("Error: Need fill LOKI_URL")
	}

	//
}

// CurrentDirectory - возвращает текущую директорию ОС
func CurrentDirectory() string {
	Otvet, err := os.Getwd()
	if err != nil {
		//log.Println(err)
	}

	return Otvet
}
