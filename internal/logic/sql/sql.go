package sql

import (
	"context"
	"fmt"
	"github.com/ManyakRus/starter/contextmain"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_pgx"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/ManyakRus/telegram_loki/internal/telegram"
	"github.com/ManyakRus/telegram_loki/internal/types"
	"github.com/jackc/pgx/v5"
	"os"
	"path"
	"strings"
	"time"
)

// Ticker - таймер, запускается каждые INTERVAL_SEND_MINUTES (10) минут
var Ticker *time.Ticker

// MapLastErrors - хранит предыдущие ошибки
var MapLastErrors = make(map[string]string)

// Start - старт работы чтения логов LOKI
func Start() {
	RunSQL_and_Send()

	Ticker = time.NewTicker(time.Duration(config.Settings.DATABASE_CHECKER_INTERVAL_MINUTES) * time.Minute)

	go ReadTicker()
}

// ReadTicker - запускается каждые INTERVAL_SEND_MINUTES минут
func ReadTicker() {
	for {
		select {
		case <-contextmain.GetContext().Done():
			return
		case <-Ticker.C:
			RunSQL_and_Send()
		}
	}

	Ticker.Stop()
}

// RunSQL_and_Send - запускает все скрипты .sql, и отправляет ошибки в Телеграм
func RunSQL_and_Send() {
	DeveloperName, err := RunSQL()
	if err != nil {
		//отправим в Telegram
		err = telegram.SendMessage(DeveloperName, err.Error())
		if err != nil {
			log.Error("SendMessage() error: ", err)
		}
	}

}

// RunSQL - запускает все скрипты .sql
// возвращает DeveloperName, error
func RunSQL() (string, error) {
	DeveloperName := ""
	var err error

	log.Debug("Start run .sql scripts")

	dir := micro.ProgramDir_bin()
	DirFilename := dir + config.Settings.SQL_FILES_FOLDER
	Files, err := os.ReadDir(DirFilename)
	if err != nil {
		err = fmt.Errorf("os.ReadDir() error: %w", err)
		log.Error(err)
		return DeveloperName, err
	}

	//пройдемся по всем файлам
	for _, file1 := range Files {
		//пропускаем файлы, которые не .sql
		//FileInfo, err := file1.Info()
		//if err != nil {
		//	err = fmt.Errorf("file1.Info() error: %w", err)
		//	log.Error(err)
		//	return err
		//}

		FilenameShort := path.Base(file1.Name())
		Filename := DirFilename + micro.SeparatorFile() + FilenameShort
		if file1.Type().IsRegular() == false || strings.HasSuffix(FilenameShort, ".sql") == false {
			continue
		}

		//запускаем скрипт
		err = RunSQL1(Filename)
		if err != nil {
			DeveloperName, _ = types.MapSQLDeveloper[FilenameShort]
			DeveloperNameTrim := FindDeveloperName_if_err(FilenameShort, err)
			err = fmt.Errorf("%w\n%s", err, DeveloperNameTrim)
			log.Warn(err)
			return DeveloperNameTrim, err
		}
	}

	return "", err
}

// FindDeveloperName_if_err - возвращает имя разработчика, если ошибка другая
func FindDeveloperName_if_err(FilenemeShort string, err error) string {
	//
	//DeveloperName := ""
	DeveloperName, ok := types.MapSQLDeveloper[FilenemeShort]
	if ok {
		//DeveloperName = DeveloperName + Name1

		//если такая же ошибка то не пишем имя разработчика
		LastError, IsFind2 := MapLastErrors[FilenemeShort]
		if IsFind2 == true && LastError == err.Error() {
			DeveloperName = ""
		}
	}

	return DeveloperName
}

// RunSQL1 - запускает 1 скрипт
func RunSQL1(Filename string) error {
	var err error

	//подключимся к БД
	db := postgres_pgx.GetConnection()

	ctxMain := contextmain.GetContext()
	ctx, cancel := context.WithTimeout(ctxMain, 10*time.Minute)
	defer cancel()

	//прочитаем файл
	bytes, err := os.ReadFile(Filename)
	if err != nil {
		err = fmt.Errorf("ReadFile() Filename: %s error: %w", Filename, err)
		log.Error(err)

		return err
	}
	TextSQL := string(bytes)

	//запустим запрос
	ResultSQL := ""
	rows := db.QueryRow(ctx, TextSQL)
	err = rows.Scan(&ResultSQL)

	//запомним последнюю ошибку
	FilenameShort := path.Base(Filename)
	TextError := ""
	if err != nil {
		TextError = err.Error()
	}
	MapLastErrors[FilenameShort] = TextError

	//нет строк - это хорошо
	if err == pgx.ErrNoRows {
		err = nil
		return err
	}

	//ошибки не должно быть
	if err != nil {
		err = fmt.Errorf("db.Exec() Filename: %s, error: %w", Filename, err)
		log.Error(err)

		return err
	}

	//запрос вернул число(строку)
	err = fmt.Errorf(`скрипт "%s" вернул значение: %s`, FilenameShort, ResultSQL)

	return err
}
