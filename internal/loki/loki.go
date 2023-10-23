package loki

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/ManyakRus/telegram_loki/internal/types"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var Client = &http.Client{}

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

	Filter := config.Settings.LOKI_SEARCH_TEXT
	URL := QueryApp(ServiceName, DateFrom, DateTo, Filter)
	//query := "%7Bapp%3D%22" + ServiceName + "%22%7D%7C~%22(error:%7Cpanic:%7CERROR:%7CPANIC:)%22"
	//sTime1 := strconv.FormatInt(DateFrom.UnixNano(), 10)
	//sTime2 := strconv.FormatInt(DateTo.UnixNano(), 10)
	//URL := config.Settings.LOKI_URL + "/api/datasources/proxy/1/loki/api/v1/query_range?direction=BACKWARD&limit=10&query=" + query
	//URL += "&start=" + sTime1 + "&end=" + sTime2

	//URL = "http://logmon.dev.atomsbt.ru/api/datasources" //удалить

	//запрос http
	req, err := http.NewRequest(http.MethodGet, URL, http.NoBody)
	if err != nil {
		return Otvet, fmt.Errorf("request(), ошибка создания GET запроса, err=%w", err)
	}

	req.SetBasicAuth(config.Settings.GRAFANA_LOGIN, config.Settings.GRAFANA_PASSWORD)
	//req.Header.Add("Content-Type:", "application/json")

	//client := &http.Client{}
	response, err := Client.Do(req)
	if err != nil {
		return Otvet, fmt.Errorf("request(), ошибка выполнения GET запроса, err=%w", err)
	}

	switch response.StatusCode {
	case 200:
	default:
		{
			TextError := fmt.Sprint("http request error: ", response.Status)
			err = errors.New(TextError)
			log.Error(TextError)
			return Otvet, err
		}
	}

	//
	TextJson, err := io.ReadAll(response.Body)
	err = json.Unmarshal(TextJson, &Otvet)
	if err != nil {
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
