package loki

import (
	"encoding/json"
	"fmt"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/ManyakRus/telegram_loki/internal/types"
	"io"
	"net/http"
	"strconv"
	"time"
)

func DownloadJSON(ServiceName string, DateFrom, DateTo time.Time) ([]types.Message, error) {
	Otvet := make([]types.Message, 0)
	var err error

	query := "%7Bapp%3D%22" + ServiceName + "%22%7D%20%7C%3D%20%22error%22"
	sTime1 := strconv.FormatInt(DateFrom.Unix(), 10)
	sTime2 := strconv.FormatInt(DateTo.Unix(), 10)
	URL := config.Settings.LOKI_URL + "/api/datasources/proxy/1/loki/api/v1/query_range?direction=BACKWARD&limit=10&query=" + query
	URL += "&start=" + sTime1 + "&end=" + sTime2

	//запрос http
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return Otvet, fmt.Errorf("request(), ошибка создания GET запроса, err=%w", err)
	}

	//req.Header.Add("Content-Type:", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return Otvet, fmt.Errorf("request(), ошибка выполнения GET запроса, err=%w", err)
	}

	switch response.StatusCode {
	case 200:
	default:
		{
			log.Error("http request error: ", response.Status)
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

func Authentication() error {
	var err error

	URL := config.Settings.LOKI_URL + "/login"

	client := http.Client{Timeout: 60 * time.Second}

	req, err := http.NewRequest(http.MethodGet, URL, http.NoBody)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(config.Settings.LOKI_LOGIN, config.Settings.LOKI_PASSWORD)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	//resBody, err := io.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//
	return err
}
