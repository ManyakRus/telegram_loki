package victoria_http

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
	"strings"
	"time"
)

// VictoriaAPI_struct - структура для работы с LOKI
type VictoriaAPI_struct struct {
}

// LayoutDateTimeRus - формат даты
const LayoutDateTimeRus = "02.01.2006T15:04:05"

// FindURL - формирует URL для запроса
func FindURL(ServiceName string, DateFrom, DateTo time.Time, Filter string) string {
	Otvet := ""

	//ServiceName = "monitor-service" //удалить

	sLimit := strconv.Itoa(config.Settings.TELEGRAM_MESSAGES_COUNT)
	Filter = url.QueryEscape(Filter)
	//sTime1 := DateFrom.Format(LayoutDateTimeRus)
	//sTime1 = sTime1 + "%2B03:00"
	//sTime2 := DateTo.Format(LayoutDateTimeRus)
	//sTime2 = sTime2 + "%2B03:00"
	sTime1 := strconv.FormatInt(DateFrom.UnixNano(), 10)
	sTime2 := strconv.FormatInt(DateTo.UnixNano(), 10)
	query := "query={kubernetes.container_name=%22" + ServiceName + "%22}AND%20~%22" + Filter + "%22&limit=" + sLimit + "&start=" + sTime1 + "&end=" + sTime2

	Otvet = config.Settings.VICTORIA_METRICS_URL + "/select/logsql/query?" + query

	return Otvet
}

// DownloadLogs - загружает логи с сайта, и возвращает массив логов
func (v VictoriaAPI_struct) DownloadLogs(ServiceName string, DateFrom, DateTo time.Time) ([]types.MessageLog, error) {
	Otvet := make([]types.MessageLog, 0)
	var err error

	MassFull, err := DownloadLogs_full(ServiceName, DateFrom, DateTo)
	if err != nil {
		err = fmt.Errorf("DownloadLogs_full(), error: %w", err)
		log.Error(err)
		return Otvet, err
	}

	//
	for _, Log1 := range MassFull {

		Date := Log1.Time
		TextLog := Log1.Msg
		if TextLog == "" {
			continue
		}
		if Date.IsZero() {
			continue
		}
		MessageLog1 := types.MessageLog{}
		MessageLog1.Date = Date
		MessageLog1.Text = TextLog

		Otvet = append(Otvet, MessageLog1)
	}

	return Otvet, err
}

// DownloadLogs_full - загружает логи с сайта, и возвращает полные логи
func DownloadLogs_full(ServiceName string, DateFrom, DateTo time.Time) ([]types.MessageVictoriaMetrics, error) {
	Otvet := make([]types.MessageVictoriaMetrics, 0)
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

	MassStrings := splitJSONStrings(string(TextJson))
	for _, Line1 := range MassStrings {
		if Line1 == "" {
			continue
		}
		VictoriaLog1 := types.MessageVictoriaMetrics{}
		err = json.Unmarshal([]byte(Line1), &VictoriaLog1)
		if err != nil {
			err = fmt.Errorf("json.Unmarshal(), ошибка десериализации JSON, URL: %s, error: %w", URL, err)
			log.Error(err)
			return Otvet, err
		}

		Otvet = append(Otvet, VictoriaLog1)
	}

	//err = json.Unmarshal(TextJson, &Otvet)
	//if err != nil {
	//	err = fmt.Errorf("json.Unmarshal(), ошибка десериализации JSON, URL: %s, error: %w", URL, err)
	//	log.Error(err)
	//	return Otvet, err
	//}

	return Otvet, err
}

// FindGrafanaURL - находит URL ссылку в LOKI на которую можно кликнуть в телеграмме
// kubernetes.container_name = ServiceName
func (v VictoriaAPI_struct) FindGrafanaURL(ServiceName string, DateFrom, DateTo time.Time) string {
	sTimeFrom := strconv.FormatInt(DateFrom.UnixMilli(), 10)
	sTimeTo := strconv.FormatInt(DateTo.UnixMilli(), 10)
	Otvet := config.Settings.GRAFANA_URL + "/explore?orgId=1&left=%5B%22" + sTimeFrom + "%22,%22" + sTimeTo + "%22,%22Loki%22,%7B%22refId%22:%22A%22,%22expr%22:%22%7Bkubernetes.container_name%3D%5C%22" + ServiceName + "%5C%22%7D%22%7D%5D"

	return Otvet
}

// splitJSONStrings - разделяет JSON на отдельные строки
func splitJSONStrings(input string) []string {
	// Разделяем строку по "}\n{"
	parts := strings.Split(input, "}\n{")

	// Восстанавливаем скобки в каждом элементе, кроме первого и последнего
	for i := range parts {
		if i > 0 {
			parts[i] = "{" + parts[i]
		}
		if i < len(parts)-1 {
			parts[i] = parts[i] + "}"
		}
	}

	return parts
}
