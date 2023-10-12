package main

import (
	ConfigMain "github.com/ManyakRus/starter/config"
	"github.com/ManyakRus/starter/stopapp"
	"github.com/ManyakRus/starter/telegram_client"
	"github.com/ManyakRus/telegram_loki/internal/config"
)

func main() {
	StartApp()
}

func StartApp() {
	ConfigMain.LoadEnv()
	config.FillSettings()

	stopapp.StartWaitStop()

	telegram_client.CreateTelegramClient(nil)
	telegram_client.ConnectTelegram()

	stopapp.GetWaitGroup_Main().Wait()
}
