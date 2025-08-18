package loki_http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/http_connect"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/ManyakRus/telegram_loki/internal/types"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// LokiAPI_struct - структура для работы с LOKI
type LokiAPI_struct struct {
}

// FindURL - формирует URL для запроса
func FindURL(ServiceName string, DateFrom, DateTo time.Time, Filter string) string {
	Otvet := ""

	slimit := "1000"
	query := "%7Bapp%3D%22" + ServiceName + "%22%7D"
	if Filter != "" {
		FilterURL := url.QueryEscape(Filter)
		query = query + "%7C~%22(" + FilterURL + ")%22"
		slimit = strconv.Itoa(config.Settings.TELEGRAM_MESSAGES_COUNT)
	}

	sTime1 := strconv.FormatInt(DateFrom.UnixNano(), 10)
	sTime2 := strconv.FormatInt(DateTo.UnixNano(), 10)
	Otvet = config.Settings.LOKI_URL + config.Settings.LOKI_API_PATH + "/loki/api/v1/query_range?direction=BACKWARD&limit=" + slimit + "&query=" + query
	//Otvet = config.Settings.LOKI_URL + "/api/datasources/proxy/1/loki/api/v1/query_range?direction=BACKWARD&limit=" + slimit + "&query=" + query
	Otvet += "&start=" + sTime1 + "&end=" + sTime2

	return Otvet
}

// DownloadLogs - загружает логи с сайта, и возвращает массив логов
func (l LokiAPI_struct) DownloadLogs(ServiceName string, DateFrom, DateTo time.Time) ([]types.MessageLog, error) {
	Otvet := make([]types.MessageLog, 0)
	var err error

	MessageLokiFull, err := DownloadLogs_full(ServiceName, DateFrom, DateTo)
	if err != nil {
		err = fmt.Errorf("DownloadLogs_full(), error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//
	for _, Result1 := range MessageLokiFull.Data.Result {

		//отправим URL логгера в Telegram
		if len(Result1.Values) == 0 {
			continue
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

			MessageLog1 := types.MessageLog{}
			MessageLog1.Date = Date
			MessageLog1.Text = TextLog

			Otvet = append(Otvet, MessageLog1)
		}
	}

	return Otvet, err
}

// DownloadLogs_full - загружает логи с сайта, и возвращает полные логи
func DownloadLogs_full(ServiceName string, DateFrom, DateTo time.Time) (types.MessageLoki, error) {
	Otvet := types.MessageLoki{}
	var err error

	if http_connect.Client == nil {
		http_connect.CreateClient()
	}

	Filter := config.Settings.SEARCH_TEXT
	URL := FindURL(ServiceName, DateFrom, DateTo, Filter)

	//запрос http
	req, err := http.NewRequest(http.MethodGet, URL, http.NoBody)
	if err != nil {
		err = fmt.Errorf("request(), ошибка создания GET запроса, URL: %s, error: %w", URL, err)
		log.Error(err)
		return Otvet, err
	}

	req.SetBasicAuth(config.Settings.LOKI_LOGIN, config.Settings.LOKI_PASSWORD)

	//таймаут 60 секунд
	ctx, cancel_ctx := context.WithTimeout(contextmain.GetContext(), 60*time.Second)
	defer cancel_ctx()
	req = req.WithContext(ctx)

	//отправка запроса
	response, err := http_connect.Client.Do(req)
	if err != nil {
		err = fmt.Errorf("Client.Do(), ошибка выполнения GET запроса, URL: %s, error: %w", URL, err)
		log.Error(err)
		return Otvet, err
	}
	defer response.Body.Close()

	//проверка статуса ответа
	switch response.StatusCode {
	case 200:
	default:
		{
			err = fmt.Errorf("request(), ошибка выполнения GET запроса, URL: %s, status: %v, error: %w", URL, response.Status, err)
			log.Error(err)
			return Otvet, err
		}
	}

	//
	TextJson, err := io.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("io.ReadAll(), ошибка чтения тела ответа, URL: %s, error: %w", URL, err)
		log.Error(err)
		return Otvet, err
	}

	err = json.Unmarshal(TextJson, &Otvet)
	if err != nil {
		err = fmt.Errorf("json.Unmarshal(), ошибка десериализации JSON, URL: %s, error: %w", URL, err)
		log.Error(err)
		return Otvet, err
	}

	return Otvet, err
}

// FindGrafanaURL - находит URL ссылку в LOKI на которую можно кликнуть в телеграмме
func (l LokiAPI_struct) FindGrafanaURL(ServiceName string, DateFrom, DateTo time.Time) string {
	sTimeFrom := strconv.FormatInt(DateFrom.UnixMilli(), 10)
	sTimeTo := strconv.FormatInt(DateTo.UnixMilli(), 10)
	Otvet := config.Settings.GRAFANA_URL + "/explore?orgId=1&left=%5B%22" + sTimeFrom + "%22,%22" + sTimeTo + "%22,%22Loki%22,%7B%22refId%22:%22A%22,%22expr%22:%22%7Bapp%3D%5C%22" + ServiceName + "%5C%22%7D%22%7D%5D"

	return Otvet
}
