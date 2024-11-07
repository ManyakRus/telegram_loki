package load_json

import (
	"encoding/json"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/telegram_loki/internal/config"
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
