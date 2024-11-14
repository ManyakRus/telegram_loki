package loki_http

import (
	ConfigMain "github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/golang-module/carbon/v2"
	"testing"
	"time"
)

func TestDownloadJSON(t *testing.T) {
	ConfigMain.LoadEnv()
	config.FillSettings()

	date2 := time.Now()
	date1 := carbon.Time2Carbon(date2).AddDays(-2).Carbon2Time()

	Otvet, err := DownloadJSON("sync-service", date1, date2)
	if err != nil {
		t.Error("TestDownloadJSON() error: ", err)
	}
	if (Otvet.Status) == "" {
		t.Error("TestDownloadJSON() error: len(Otvet) =0")
	}
}

func TestAuthentication(t *testing.T) {
	ConfigMain.LoadEnv()
	config.FillSettings()

	err := Authentication()
	if err != nil {
		t.Error("TestAuthentication() error: ", err)
	}
}
