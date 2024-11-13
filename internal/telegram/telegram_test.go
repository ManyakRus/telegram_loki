package telegram

import (
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/telegram_bot"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/ManyakRus/telegram_loki/internal/load_json"
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	config_main.LoadEnvTest()
	config.FillSettings()
	load_json.LoadJSON_All()

	telegram_bot.Connect()
	defer telegram_bot.CloseConnection()

	Text := "test " + time.Now().String()
	SendMessage("", Text)
}
