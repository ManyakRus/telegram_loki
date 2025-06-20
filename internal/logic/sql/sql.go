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
	RunSQL_all()

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
			RunSQL_all()
		}
	}

	Ticker.Stop()
}

// RunSQL_all_and_Send - запускает все скрипты .sql, и отправляет ошибки в Телеграм
//func RunSQL_all_and_Send() {
//	Message1, err := RunSQL_all()
//
//}

// RunSQL_all - запускает все скрипты .sql
// возвращает DeveloperName, error
func RunSQL_all() {
	Otvet := types.Message{}
	var err error

	log.Debug("Start run .sql scripts")

	dir := micro.ProgramDir_bin()
	DirFilename := dir + config.Settings.SQL_FILES_FOLDER
	Files, err := os.ReadDir(DirFilename)
	if err != nil {
		err = fmt.Errorf("os.ReadDir() error: %w", err)
		log.Error(err)
		Otvet.Text = err.Error()

		//отправим в телеграм
		err = telegram.SendMessage(Otvet)
		if err != nil {
			log.Error("SendMessage() error: ", err)
		}

		return
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

		Otvet = types.Message{}
		//Otvet.ServiceName = FilenameShort

		//запускаем скрипт
		Text := RunSQL1(Filename)
		if Text != "" {
			//проверим что ошибка такая же
			IsSameLastText := false
			LastText, _ := MapLastErrors[FilenameShort]
			if Text == LastText {
				IsSameLastText = true
			}

			//
			DeveloperName, _ := types.MapSQLDeveloper[FilenameShort]
			Otvet.DeveloperName = DeveloperName
			Otvet.Text = Text
			Otvet.IsSameLastText = IsSameLastText

			//DeveloperNameTrim := FindDeveloperName_if_err(FilenameShort, err)
			//err2 := fmt.Errorf("%w\n%s", err, DeveloperNameTrim)
			//log.Warn(err2)

			//запомним ошибку
			MapLastErrors[FilenameShort] = Text

			//отправим в телеграм
			if Otvet.Text != "" {
				err = telegram.SendMessage(Otvet)
				if err != nil {
					log.Error("SendMessage() error: ", err)
				}
			}

		}
	}

	return
}

// FindDeveloperName_if_err - возвращает имя разработчика, если ошибка другая
func FindDeveloperName_if_err(FilenemeShort string, err error) string {
	//
	if err == nil {
		return ""
	}

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
func RunSQL1(Filename string) string {
	Otvet := ""
	var err error

	//прочитаем файл
	bytes, err := os.ReadFile(Filename)
	if err != nil {
		Otvet = fmt.Sprintf("ReadFile() Filename: %s error: %v", Filename, err)
		log.Error(Otvet)

		return Otvet
	}
	TextSQL := string(bytes)

	//подключимся к БД
	db := postgres_pgx.GetConnection()
	ctxMain := contextmain.GetContext()
	ctx, cancel := context.WithTimeout(ctxMain, 10*time.Minute)
	defer cancel()

	//запустим запрос
	ResultSQL := ""
	rows := db.QueryRow(ctx, TextSQL)
	err = rows.Scan(&ResultSQL)

	//запомним последнюю ошибку
	FilenameShort := path.Base(Filename)

	//нет строк - это хорошо
	if err == pgx.ErrNoRows {
		return Otvet
	}

	//ошибки не должно быть
	if err != nil {
		Otvet = fmt.Sprintf("db.QueryRow() Filename: %s, error: %v", FilenameShort, err)
		log.Error(Otvet)
		return Otvet
	}

	//запрос вернул число(строку)
	sDate := micro.StringDateTime(time.Now())
	Otvet = fmt.Sprintf(`скрипт "%s" вернул значение: %s`+"\n"+`Дата: %s`, FilenameShort, ResultSQL, sDate)

	return Otvet
}
