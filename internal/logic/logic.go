package logic

import (
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/ManyakRus/telegram_loki/internal/logic/loki"
	"github.com/ManyakRus/telegram_loki/internal/logic/sql"
)

// Start - старт работы логики программы
func Start() {
	//LOKI
	if config.Settings.LOKI_CHECKER_ENABLED == true {
		loki.Start()
	}

	//SQL
	if config.Settings.DATABASE_CHECKER_ENABLED == true {
		sql.Start()
	}
}
