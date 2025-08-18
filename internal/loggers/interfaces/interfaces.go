package interfaces

import (
	"github.com/ManyakRus/telegram_loki/internal/types"
	"time"
)

// ILogger - интерфейс для разных логгеров
type ILogger interface {
	DownloadLogs(ServiceName string, DateFrom, DateTo time.Time) ([]types.MessageLog, error)
	FindURL(ServiceName string, DateFrom, DateTo time.Time, Filter string) string
}
