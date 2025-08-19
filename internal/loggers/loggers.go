package loggers

import (
	"github.com/ManyakRus/telegram_loki/internal/config"
	"strconv"
	"time"
)

// FindGrafanaURL - находит URL ссылку в LOKI на которую можно кликнуть в телеграмме
func FindGrafanaURL(ServiceName string, DateFrom, DateTo time.Time) string {
	sTimeFrom := strconv.FormatInt(DateFrom.UnixMilli(), 10)
	sTimeTo := strconv.FormatInt(DateTo.UnixMilli(), 10)
	Otvet := config.Settings.GRAFANA_URL + "/explore?orgId=1&left=%5B%22" + sTimeFrom + "%22,%22" + sTimeTo + "%22,%22Loki%22,%7B%22refId%22:%22A%22,%22expr%22:%22%7Bkubernetes.container_name%3D%5C%22" + ServiceName + "%5C%22%7D%22%7D%5D"

	return Otvet
}
