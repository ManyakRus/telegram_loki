package config

import (
	"log"
	"os"
	"strconv"
)

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	GRAFANA_LOGIN           string
	GRAFANA_PASSWORD        string
	LOKI_URL                string
	LOKI_SEARCH_TEXT        string
	TELEGRAM_CHAT_NAME      string
	TELEGRAM_MESSAGES_COUNT int
	INTERVAL_SEND_MINUTES   int
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {
	var err error

	Settings = SettingsINI{}
	Settings.TELEGRAM_CHAT_NAME = os.Getenv("TELEGRAM_CHAT_NAME")
	Settings.LOKI_URL = os.Getenv("LOKI_URL")
	Settings.GRAFANA_LOGIN = os.Getenv("GRAFANA_LOGIN")
	Settings.GRAFANA_PASSWORD = os.Getenv("GRAFANA_PASSWORD")
	Settings.INTERVAL_SEND_MINUTES, err = strconv.Atoi(os.Getenv("INTERVAL_SEND_MINUTES"))
	if err != nil {
		Settings.INTERVAL_SEND_MINUTES = 10
	}
	Settings.TELEGRAM_MESSAGES_COUNT, err = strconv.Atoi(os.Getenv("TELEGRAM_MESSAGES_COUNT"))
	if err != nil || Settings.TELEGRAM_MESSAGES_COUNT == 0 {
		Settings.TELEGRAM_MESSAGES_COUNT = 1
	}
	s := os.Getenv("LOKI_SEARCH_TEXT")
	Settings.LOKI_SEARCH_TEXT = s
	if Settings.LOKI_SEARCH_TEXT == "" {
		Settings.LOKI_SEARCH_TEXT = "error:|panic:|ERROR:|PANIC:"
		//Settings.LOKI_SEARCH_TEXT = "error:%7Cpanic:%7CERROR:%7CPANIC:"
	}

	if Settings.TELEGRAM_CHAT_NAME == "" {
		log.Panic("Error: Need fill TELEGRAM_CHAT_NAME")
	}

	if Settings.LOKI_URL == "" {
		log.Panic("Error: Need fill LOKI_URL")
	}

	if Settings.GRAFANA_LOGIN == "" {
		log.Panic("Error: Need fill GRAFANA_LOGIN")
	}

	if Settings.GRAFANA_PASSWORD == "" {
		log.Panic("Error: Need fill GRAFANA_PASSWORD")
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
