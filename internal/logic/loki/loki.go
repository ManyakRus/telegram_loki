package loki

import (
	"fmt"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/ManyakRus/telegram_loki/internal/loggers/interfaces"
	"github.com/ManyakRus/telegram_loki/internal/loggers/loki_http"
	"github.com/ManyakRus/telegram_loki/internal/telegram"
	"github.com/ManyakRus/telegram_loki/internal/types"
	"github.com/dromara/carbon/v2"
	"time"
)

// LastReadTime - время последнего чтения логов
var LastReadTime time.Time

// Ticker - таймер, запускается каждые INTERVAL_SEND_MINUTES (10) минут
var Ticker *time.Ticker

// MapLastErrors - хранит предыдущие ошибки
var MapLastErrors = make(map[string]string)

// Start - старт работы чтения логов LOKI
func Start() {
	var err error

	var LoggerAPI interfaces.ILogger
	if config.Settings.LOKI_URL != "" {
		LoggerAPI = loki_http.LokiAPI_struct{}
	} else if config.Settings.VICTORIA_METRICS_URL != "" {
		//LoggerAPI = victoria_http.VictoriaAPI{}
	} else {
		err = fmt.Errorf("no LOKI_URL or VICTORIA_METRICS_URL")
		log.Error(err)
		log.Panic(err)
	}

	//
	log.Info("Start read LOKI, URL: ", config.Settings.LOKI_URL+config.Settings.LOKI_API_PATH)

	//
	Date2 := time.Now()
	Date1 := carbon.CreateFromStdTime(Date2).AddMinutes(-1 * config.Settings.LOKI_CHECKER_INTERVAL_MINUTES).StdTime()
	go Start_period(LoggerAPI, Date1, Date2)

	Ticker = time.NewTicker(time.Duration(config.Settings.LOKI_CHECKER_INTERVAL_MINUTES) * time.Minute)
	//defer Ticker.Stop()

	go ReadTicker(LoggerAPI)
}

// ReadTicker - запускается каждые INTERVAL_SEND_MINUTES минут
func ReadTicker(LoggerAPI interfaces.ILogger) {
	for {
		select {
		case <-contextmain.GetContext().Done():
			return
		case <-Ticker.C:
			Date2 := time.Now()
			Date1 := LastReadTime
			Start_period(LoggerAPI, Date1, Date2)
		}
	}

	Ticker.Stop()
}

// Start_period - запускает чтение логов всех сервисов за период
func Start_period(LoggerAPI interfaces.ILogger, Date1, Date2 time.Time) {
	sDate1 := micro.StringDateTime(Date1)
	sDate2 := micro.StringDateTime(Date2)
	log.Debug("Start search errors from: ", sDate1, " to: ", sDate2)
	IsOnlyErrors := true
	var err1 error
	var err error

	ctxMain := contextmain.GetContext()

	CountText := 0
loop_for:
	for ServiceName, DeveloperName := range types.MapServiceDeveloper {
		//проверка завершения программы
		select {
		case <-contextmain.GetContext().Done():
			log.Warn("Context app is canceled. Start()")
			break loop_for
		default:
		}

		//запуск проверки логов одного сервиса
		Message1 := types.Message{}
		Message1.ServiceName = ServiceName
		Message1.DeveloperName = DeveloperName
		err1 = Start_period1(LoggerAPI, &Message1, Date1, Date2)
		if err1 != nil {
			micro.Pause_ctx(ctxMain, 1000) //error: 502 Bad Gateway
		} else {
			if Message1.Text != "" {
				CountText = CountText + 1
			}
			IsOnlyErrors = false
			micro.Pause_ctx(ctxMain, 100) //error: 429 Too Many Requests
		}

	}

	LastReadTime = Date2

	if CountText == 0 {
		log.Debug(`No text "error" found in LOKI`)
	}

	// если только ошибки - то напишем в телеграм
	if IsOnlyErrors == true {
		TextError := fmt.Sprint("Search errors: only errors. Last error: ", err1)
		log.Debug(TextError)
		Message1 := types.Message{}
		Message1.Text = TextError
		err = telegram.SendMessage(Message1)
		if err != nil {
			log.Error("telegram.SendMessage() error: ", err)
		}
		micro.Pause_ctx(ctxMain, 60*60*1000) //пауза 1 час
	}
}

// Start_period1 - запускает чтение логов одного сервиса за период
// возвращает Текст отправленного сообщения, и ошибку
func Start_period1(LoggerAPI interfaces.ILogger, Message1 *types.Message, DateFrom, DateTo time.Time) error {
	var err error
	//Text := ""

	ctxMain := contextmain.GetContext()
	DeveloperName0 := Message1.DeveloperName

	MassMessageLog, err := LoggerAPI.DownloadLogs(Message1.ServiceName, DateFrom, DateTo)
	if err != nil {
		//log.Error("DownloadLogs() error: ", err)
		micro.Pause_ctx(ctxMain, 1000)
		return err
	}

	for _, Log1 := range MassMessageLog {

		//
		URL := LoggerAPI.FindURL(Message1.ServiceName, DateFrom, DateTo)

		TextLog := Log1.Text
		Date := Log1.Date

		DeveloperName := DeveloperName0

		//если такая же ошибка то не пишем имя разработчика
		IsSameLastText := false
		LastText := ""
		LastText, _ = MapLastErrors[Message1.ServiceName]
		TextLogWithoutTime := FindTextWithoutTime(TextLog)
		TextLastLogWithoutTime := FindTextWithoutTime(LastText)
		if TextLogWithoutTime == TextLastLogWithoutTime {
			IsSameLastText = true
		}

		//запомним последнюю ошибку
		MapLastErrors[Message1.ServiceName] = TextLog

		//
		//Text = TextServiceName + " " + TextDate + DeveloperName + "\n" + TextLog
		//Text = TextLog
		Message1.Text = TextLog
		Message1.DeveloperName = DeveloperName
		Message1.Date = Date
		Message1.LokiURL = URL
		Message1.IsSameLastText = IsSameLastText

		//
		err = telegram.SendMessage(*Message1)
		if err != nil {
			log.Error("SendMessage() error: ", err)
			continue
		}
	}

	return err
}

//// FindURLLoki - находит URL ссылку в LOKI на которую можно кликнуть в телеграмме
//func FindURLLoki(ServiceName string, DateFrom, DateTo time.Time) string {
//	sTimeFrom := strconv.FormatInt(DateFrom.UnixMilli(), 10)
//	sTimeTo := strconv.FormatInt(DateTo.UnixMilli(), 10)
//	Otvet := config.Settings.LOKI_URL + "/explore?orgId=1&left=%5B%22" + sTimeFrom + "%22,%22" + sTimeTo + "%22,%22Loki%22,%7B%22refId%22:%22A%22,%22expr%22:%22%7Bapp%3D%5C%22" + ServiceName + "%5C%22%7D%22%7D%5D"
//
//	return Otvet
//}

// FindTextWithoutTime - убирает время в логе, для этого берём текст после 2 пробела
// time="2024-11-07 04:30:53.709" level=error msg="GetExtractEgripEgrul INN: 519300250706 Result: map[address: fullName: shortName: state:] Error: 404 Нет данных по данному ИНН: 519300250706" func="TakeMessageAsync()\t" file=" nats.go:194\t"
// 2024/11/01 04:35:07.872721 [ERROR] syncMessage, contractId: 7802, error: NewBriefCase, initInvoices, error: there is no documents for this contract
func FindTextWithoutTime(TextLog string) string {
	Otvet := TextLog

	Otvet = micro.StringAfter(Otvet, " ")
	Otvet = micro.StringAfter(Otvet, " ")

	return Otvet
}
