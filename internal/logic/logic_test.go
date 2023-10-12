package logic

import (
	ConfigMain "github.com/ManyakRus/starter/config"
	"github.com/ManyakRus/starter/telegram_client"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/ManyakRus/telegram_loki/internal/constants"
	"github.com/ManyakRus/telegram_loki/internal/load_json"
	"github.com/golang-module/carbon/v2"
	"testing"
	"time"
)

func TestStart_period(t *testing.T) {
	ConfigMain.LoadEnv()
	config.FillSettings()

	telegram_client.CreateTelegramClient(nil)
	telegram_client.ConnectTelegram()
	load_json.LoadJSON()

	date2 := time.Now()
	date1 := carbon.Time2Carbon(date2).AddDays(-2).Carbon2Time()

	ServiceName := "sync-service"
	DeveloperName := "@ManyakRus"
	Start_period1(ServiceName, DeveloperName, date1, date2)
}

func TestTime(t *testing.T) {
	TextDate := time.Now().Format(constants.Layout)
	//TextDate := fmt.Sprintf(constants.Layout, time.Now())
	print(TextDate)
}
