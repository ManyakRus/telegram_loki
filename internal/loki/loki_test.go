package loki

import (
	ConfigMain "github.com/ManyakRus/starter/config"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/golang-module/carbon/v2"
	"testing"
	"time"
)

func TestDownloadJSON(t *testing.T) {
	ConfigMain.LoadEnv()
	config.FillSettings()

	date1 := time.Now()
	date2 := carbon.Time2Carbon(date1).AddDays(-1).Carbon2Time()

	Otvet, err := DownloadJSON("sync-service", date1, date2)
	if err != nil {
		t.Error("TestDownloadJSON() error: ", err)
	}
	if len(Otvet) == 0 {
		t.Error("TestDownloadJSON() error: len(Otvet) =0")
	}
}

func TestAuthentication(t *testing.T) {
	err := Authentication()
	if err != nil {
		t.Error("TestAuthentication() error: ", err)
	}
}
