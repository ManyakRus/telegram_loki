package load_json

import (
	"encoding/json"
	"fmt"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"github.com/ManyakRus/telegram_loki/internal/constants"
	"github.com/ManyakRus/telegram_loki/internal/types"
	"os"
)

// LoadJSON_All - загружает 2 файла в формате .json в map
func LoadJSON_All() {
	var err error
	dir := micro.ProgramDir_bin()

	log.Info("app directory: ", dir)

	//services.txt
	FileName := dir + "settings" + micro.SeparatorFile() + "services.txt"
	err = LoadJSON_Services(FileName)
	if err != nil {
		log.Panic("LoadJSON_Services() error: ", err)
	}

	//services_add.txt, дополнительный файл, необязательный
	FileName = dir + "settings" + micro.SeparatorFile() + "services_add.txt"
	ok, err := micro.FileExists(FileName)
	if ok == true {
		err = LoadJSON_Services(FileName)
		if err != nil {
			log.Debug("LoadJSON_Services() warning: ", err)
		}
	}

	//scripts.txt
	FileName = dir + config.Settings.SQL_FILES_FOLDER + micro.SeparatorFile() + "scripts.txt"
	err = LoadJSON_SQL(FileName)
	if err != nil {
		log.Panic("LoadJSON_Services() error: ", err)
	}

	//scripts_add.txt, дополнительный файл, необязательный
	FileName = dir + config.Settings.SQL_FILES_FOLDER + micro.SeparatorFile() + "scripts_add.txt"
	ok, err = micro.FileExists(FileName)
	if ok == true {
		err = LoadJSON_SQL(FileName)
		if err != nil {
			log.Debug("LoadJSON_Services() warning: ", err)
		}
	}

	//telegram_users.json
	FileName = dir + config.Settings.TELEGRAM_USERS_FILENAME
	ok, err = micro.FileExists(FileName)
	if ok == true {
		err = LoadJSON_TelegramUsers(FileName)
		if err != nil {
			log.Debug("LoadJSON_TelegramUsers() warning: ", err)
		}
	}
}

// LoadJSON_Services - загружает файл в формате .json в map, services.txt и services_add.txt
func LoadJSON_Services(FileName string) error {
	var err error

	//чтение файла
	bytes, err := os.ReadFile(FileName)
	if err != nil {
		log.Error("ReadFile() error: ", err)
		return err
	}

	//json в map
	var MapServiceURL2 = make(map[string]string)
	err = json.Unmarshal(bytes, &MapServiceURL2)
	if err != nil {
		log.Panic("Unmarshal() error: ", err)
	}

	//заполнение главного map
	for k, v := range MapServiceURL2 {
		types.MapServiceDeveloper[k] = v
	}

	return err
}

// LoadJSON_SQL - загружает файл в формате .json в map, services.txt и services_add.txt
func LoadJSON_SQL(FileName string) error {
	var err error

	//чтение файла
	bytes, err := os.ReadFile(FileName)
	if err != nil {
		log.Error("ReadFile() error: ", err)
		return err
	}

	//json в map
	var MapServiceURL2 = make(map[string]string)
	err = json.Unmarshal(bytes, &MapServiceURL2)
	if err != nil {
		log.Panic("Unmarshal() error: ", err)
	}

	//заполнение главного map
	for k, v := range MapServiceURL2 {
		types.MapSQLDeveloper[k] = v
	}

	return err
}

// LoadJSON_TelegramUsers - загружает файл в формате .json в map, telegram_users.json
func LoadJSON_TelegramUsers(FileName string) error {
	var err error

	//чтение файла
	bytes, err := os.ReadFile(FileName)
	if err != nil {
		log.Error("ReadFile() error: ", err)
		return err
	}

	//json в map
	err = json.Unmarshal(bytes, &types.MapTelegramUsers)
	if err != nil {
		log.Panic("Unmarshal() error: ", err)
	}

	return err
}

// SaveMapTelegramUsers - записывает map в файл
func SaveMapTelegramUsers() {
	var err error
	err = SaveMapTelegramUsers_err()
	if err != nil {
		log.Error("SaveMapTelegramUsers_err() error: ", err)
	}
}

// SaveMapTelegramUsers_err - записывает map в файл
func SaveMapTelegramUsers_err() error {
	var err error

	dir := micro.ProgramDir_bin()
	Filename := dir + config.Settings.TELEGRAM_USERS_FILENAME
	bytes, err := json.MarshalIndent(types.MapTelegramUsers, "", " ")
	if err != nil {
		err = fmt.Errorf("MarshalIndent() error: %w", err)
		return err
	}

	//запись в файл
	err = os.WriteFile(Filename, bytes, constants.FILE_PERMISSIONS)
	if err != nil {
		err = fmt.Errorf("WriteFile() error: %w", err)
		return err
	}

	return err
}
