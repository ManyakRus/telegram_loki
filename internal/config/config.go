package config

import (
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/microl"
	"os"
	"strconv"
)

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	GRAFANA_LOGIN                     string
	GRAFANA_PASSWORD                  string
	LOKI_URL                          string
	LOKI_SEARCH_TEXT                  string
	TELEGRAM_CHAT_NAME                string
	TELEGRAM_MESSAGES_COUNT           int
	LOKI_CHECKER_INTERVAL_MINUTES     int
	LOKI_CHECKER_ENABLED              bool
	DATABASE_CHECKER_ENABLED          bool
	DATABASE_CHECKER_INTERVAL_MINUTES int
	SQL_FILES_FOLDER                  string
	TELEGRAM_USERS_FILENAME           string
	LOKI_API_PATH                     string
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {
	var err error

	Name := ""

	Settings = SettingsINI{}
	Settings.TELEGRAM_CHAT_NAME = os.Getenv("TELEGRAM_CHAT_NAME")
	Settings.LOKI_URL = os.Getenv("LOKI_URL")
	Settings.GRAFANA_LOGIN = os.Getenv("GRAFANA_LOGIN")
	Settings.GRAFANA_PASSWORD = os.Getenv("GRAFANA_PASSWORD")
	Settings.LOKI_CHECKER_INTERVAL_MINUTES, err = strconv.Atoi(os.Getenv("LOKI_CHECKER_INTERVAL_MINUTES"))
	if err != nil {
		Settings.LOKI_CHECKER_INTERVAL_MINUTES = 10
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
	Name = "LOKI_CHECKER_ENABLED"
	s = Getenv(Name, true)
	Settings.LOKI_CHECKER_ENABLED = micro.BoolFromString(s)

	//
	Name = "DATABASE_CHECKER_ENABLED"
	s = Getenv(Name, true)
	Settings.DATABASE_CHECKER_ENABLED = micro.BoolFromString(s)

	//
	Name = "DATABASE_CHECKER_INTERVAL_MINUTES"
	s = Getenv(Name, true)
	x, err := strconv.Atoi(s)
	if err != nil {
		x = 60
		log.Warn(Name+" error: ", err)
	}
	Settings.DATABASE_CHECKER_INTERVAL_MINUTES = x

	//
	Name = "SQL_FILES_FOLDER"
	s = Getenv(Name, true)
	Settings.SQL_FILES_FOLDER = s

	//
	Name = "TELEGRAM_USERS_FILENAME"
	s = Getenv(Name, true)
	Settings.TELEGRAM_USERS_FILENAME = s

	//
	Name = "LOKI_API_PATH"
	microl.Set_FieldFromEnv_String(&Settings, Name, true)

}

// CurrentDirectory - возвращает текущую директорию ОС
func CurrentDirectory() string {
	Otvet, err := os.Getwd()
	if err != nil {
		//log.Println(err)
	}

	return Otvet
}

// Getenv - возвращает переменную окружения
func Getenv(Name string, IsRequired bool) string {
	TextError := "Need fill OS environment variable: "
	Otvet := os.Getenv(Name)
	if IsRequired == true && Otvet == "" {
		log.Error(TextError + Name)
	}

	return Otvet
}
