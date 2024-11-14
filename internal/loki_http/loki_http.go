package loki_http

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/ManyakRus/telegram_loki/internal/types"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var Client *http.Client // &http.Client{}

func QueryApp(ServiceName string, DateFrom, DateTo time.Time, Filter string) string {
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
	Otvet = config.Settings.LOKI_URL + "/api/datasources/proxy/1/loki/api/v1/query_range?direction=BACKWARD&limit=" + slimit + "&query=" + query
	Otvet += "&start=" + sTime1 + "&end=" + sTime2

	return Otvet
}

func DownloadJSON(ServiceName string, DateFrom, DateTo time.Time) (types.Message, error) {
	Otvet := types.Message{}
	var err error

	if Client == nil {
		CreateClient()
	}

	Filter := config.Settings.LOKI_SEARCH_TEXT
	URL := QueryApp(ServiceName, DateFrom, DateTo, Filter)

	//запрос http
	req, err := http.NewRequest(http.MethodGet, URL, http.NoBody)
	if err != nil {
		err = fmt.Errorf("request(), ошибка создания GET запроса, URL: %s, error: %w", URL, err)
		log.Error(err)
		return Otvet, err
	}

	req.SetBasicAuth(config.Settings.GRAFANA_LOGIN, config.Settings.GRAFANA_PASSWORD)
	//req.Header.Add("Content-Type:", "application/json")

	//таймаут 60 секунд
	ctx, cancel_ctx := context.WithTimeout(contextmain.GetContext(), 60*time.Second)
	defer cancel_ctx()
	req = req.WithContext(ctx)

	//отправка запроса
	response, err := Client.Do(req)
	if err != nil {
		err = fmt.Errorf("Client.Do(), ошибка выполнения GET запроса, URL: %s, error: %w", URL, err)
		log.Error(err)
		return Otvet, err
	}

	//проверка статуса ответа
	switch response.StatusCode {
	case 200:
	default:
		{
			TextError := fmt.Sprintf("http request URL: %s, error status: %s", URL, response.Status)
			err = errors.New(TextError)
			log.Error(TextError)
			return Otvet, err
		}
	}

	//
	TextJson, err := io.ReadAll(response.Body)
	err = json.Unmarshal(TextJson, &Otvet)
	if err != nil {
		err = fmt.Errorf("json.Unmarshal(), ошибка десериализации JSON, URL: %s, error: %w", URL, err)
		log.Error(err)
		return Otvet, err
	}

	return Otvet, err
}

// Authentication - ненужная
// функция для аутентификации
func Authentication() error {
	var err error

	URL := config.Settings.LOKI_URL //+ "/login"

	//client := http.Client{Timeout: 60 * time.Second}

	req, err := http.NewRequest(http.MethodGet, URL, http.NoBody)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(config.Settings.GRAFANA_LOGIN, config.Settings.GRAFANA_PASSWORD)

	res, err := Client.Do(req)
	if err != nil {
		log.Panic(err)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	print(resBody)

	//
	return err
}

// CreateClient - создаёт клиент http
// для работы без ошибки сертификатов
func CreateClient() {
	Client = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig: &tls.Config{
				// See comment above.
				// UNSAFE!
				// DON'T USE IN PRODUCTION!
				InsecureSkipVerify: true,
			},
		},
	}
}
