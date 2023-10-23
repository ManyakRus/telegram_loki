package logic

import (
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/telegram_client"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/ManyakRus/telegram_loki/internal/constants"
	"github.com/ManyakRus/telegram_loki/internal/load_json"
	"github.com/ManyakRus/telegram_loki/internal/loki"
	"github.com/ManyakRus/telegram_loki/internal/types"
	"github.com/golang-module/carbon/v2"
	"strconv"
	"time"
)

//var MassTable []types.Table

var LastReadTime time.Time

var Ticker *time.Ticker

// Start - старт работы логики программы
func Start() {
	load_json.LoadJSON()

	Date2 := time.Now()
	Date1 := carbon.Time2Carbon(Date2).AddMinutes(-1 * config.Settings.INTERVAL_SEND_MINUTES).Carbon2Time()
	Start_period(Date1, Date2)

	Ticker = time.NewTicker(time.Duration(config.Settings.INTERVAL_SEND_MINUTES) * time.Minute)
	defer Ticker.Stop()

	ReadTicker()
}

// ReadTicker - запускается каждые INTERVAL_SEND_MINUTES минут
func ReadTicker() {
	for {
		select {
		case <-contextmain.GetContext().Done():
			return
		case <-Ticker.C:
			Date2 := time.Now()
			Date1 := LastReadTime
			Start_period(Date1, Date2)
		}
	}
}

// Start_period - запускает чтение логов всех сервисов за период
func Start_period(Date1, Date2 time.Time) {
	log.Debug("Start search errors from: ", Date1, " to: ", Date2)
	for ServiceName, DeveloperName := range types.MapServiceDeveloper {
		select {
		case <-contextmain.GetContext().Done():
			log.Warn("Context app is canceled. Start()")
		default:
			Start_period1(ServiceName, DeveloperName, Date1, Date2)
			micro.Pause(100) //error: 429 Too Many Requests
		}

	}
	LastReadTime = Date2
}

// Start_period1 - запускает чтение логов одного сервиса за период
func Start_period1(ServiceName, DeveloperName string, DateFrom, DateTo time.Time) error {
	var err error

	LokiMessage, err := loki.DownloadJSON(ServiceName, DateFrom, DateTo)
	if err != nil {
		log.Error("DownloadJSON() error: ", err)
		micro.Pause(1000)
		return err
	}

	for _, Result1 := range LokiMessage.Data.Result {

		//отправим URL логгера в Telegram
		if len(Result1.Values) > 0 {
			URL := FindURLLoki(ServiceName, DateFrom, DateTo)
			_, err = telegram_client.SendMessage(config.Settings.TELEGRAM_CHAT_NAME, URL)
			if err != nil {
				log.Error("SendMessage() error: ", err)
				continue
			}
		}

		//отправим ошибки в Telegram
		for _, MassValues1 := range Result1.Values {

			if len(MassValues1) != 2 {
				log.Warnf("строк !=2 MassValues: #%v", MassValues1)
				continue
			}
			sDate := MassValues1[0]
			TextLog := MassValues1[1]
			if TextLog == "" {
				continue
			}
			if sDate == "" {
				continue
			}
			iDate, err := strconv.ParseInt(sDate, 10, 64)
			if err != nil {
				log.Error("ParseInt() error: ", err)
				continue
			}
			Date := time.Unix(0, iDate)
			TextDate := Date.Format(constants.Layout)
			Text := ServiceName + " " + TextDate + " " + DeveloperName + "\n" + TextLog

			_, err = telegram_client.SendMessage(config.Settings.TELEGRAM_CHAT_NAME, Text)
			if err != nil {
				log.Error("SendMessage() error: ", err)
				continue
			}
		}
	}

	return err
}

// FindURLLoki - находит URL ссылку в LOKI на которую можно кликнуть в телеграмме
func FindURLLoki(ServiceName string, DateFrom, DateTo time.Time) string {
	sTimeFrom := strconv.FormatInt(DateFrom.UnixMilli(), 10)
	sTimeTo := strconv.FormatInt(DateTo.UnixMilli(), 10)
	Otvet := config.Settings.LOKI_URL + "/explore?orgId=1&left=%5B%22" + sTimeFrom + "%22,%22" + sTimeTo + "%22,%22Loki%22,%7B%22refId%22:%22A%22,%22expr%22:%22%7Bapp%3D%5C%22" + ServiceName + "%5C%22%7D%22%7D%5D"

	return Otvet
}
