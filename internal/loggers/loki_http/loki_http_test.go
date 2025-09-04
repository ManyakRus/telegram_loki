package loki_http

import (
	ConfigMain "github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/dromara/carbon/v2"
	"testing"
	"time"
)

func TestDownloadJSON(t *testing.T) {
	ConfigMain.LoadEnv()
	config.FillSettings()

	date2 := time.Now()
	date1 := carbon.NewCarbon(date2).AddDays(-2).StdTime()

	Otvet, err := DownloadLogs_full("sync-service", date1, date2)
	if err != nil {
		t.Error("TestDownloadJSON() error: ", err)
	}
	log.Debugf("Otvet: %+v", Otvet)
	//if (Otvet.Status) == "" {
	//	t.Error("TestDownloadJSON() error: len(Otvet) =0")
	//}
}

//func TestAuthentication(t *testing.T) {
//	ConfigMain.LoadEnv()
//	config.FillSettings()
//
//	err := Authentication()
//	if err != nil {
//		t.Error("TestAuthentication() error: ", err)
//	}
//}
