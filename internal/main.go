package main

import (
	ConfigMain "github.com/ManyakRus/starter/config"
	"github.com/ManyakRus/starter/telegram_client"
	"github.com/ManyakRus/telegram_loki/internal/config"
)

func main() {
	StartApp()
}

func StartApp() {
	ConfigMain.LoadEnv()
	config.FillSettings()

	telegram_client.ConnectTelegram()
}
