package loki

import (
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/ManyakRus/telegram_loki/internal/load_json"
	"testing"
	"time"
)

func TestStart_period(t *testing.T) {
	config_main.LoadEnvTest()
	config.FillSettings()
	load_json.LoadJSON_All()

	Date2 := time.Now()
	Date1 := carbon.CreateFromStdTime(Date2).AddMinutes(-200).StdTime()
	Start_period(Date1, Date2)

}
