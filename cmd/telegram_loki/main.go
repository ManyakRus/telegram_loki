package main

import (
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_pgx"
	"github.com/ManyakRus/starter/stopapp"
	"github.com/ManyakRus/starter/telegram_bot"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/ManyakRus/telegram_loki/internal/constants"
	"github.com/ManyakRus/telegram_loki/internal/load_json"
	"github.com/ManyakRus/telegram_loki/internal/logic"
	"github.com/ManyakRus/telegram_loki/internal/telegram"
	"github.com/ManyakRus/telegram_loki/pkg/version"
)

func main() {
	StartApp()
}

func StartApp() {
	micro.Show_Version(version.Version)

	config_main.LoadENV_or_SettingsTXT()
	config.FillSettings()
	load_json.LoadJSON_All()

	stopapp.StartWaitStop()

	if config.Settings.DATABASE_CHECKER_ENABLED == true {
		postgres_pgx.Start(constants.SERVICE_NAME)
	}

	//telegram_client.Connect(nil)

	telegram_bot.Start()

	telegram.StartReadMessages()

	logic.Start()

	stopapp.GetWaitGroup_Main().Wait()

	//
	log.Info("Application stopped.")
}
