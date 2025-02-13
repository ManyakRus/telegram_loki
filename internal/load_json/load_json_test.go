package load_json

import (
	"encoding/json"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/telegram_loki/internal/logic"
	"os"
	"testing"
)

func Test_SaveJSON_to_file(t *testing.T) {

	logic.MapServiceURL["Service1"] = "github.com/ManyakRus/1"
	logic.MapServiceURL["Service2"] = "github.com/ManyakRus/2"

	bytes, err := json.MarshalIndent(logic.MapServiceURL, "", "  ")
	if err != nil {
		t.Error("Test_SaveJSON_to_file() error: ", err)
	}

	dir := micro.ProgramDir()
	FileName := dir + "settings" + micro.SeparatorFile() + "connections_test.txt"
	File, err := os.Create(FileName)
	if err != nil {
		t.Error("Test_SaveJSON_to_file() error: ", err)
	}
	File.Write(bytes)
	File.Close()
}

func TestLoadJSON_from_file(t *testing.T) {
	dir := micro.ProgramDir()

	//главный файл
	FileName := dir + "settings" + micro.SeparatorFile() + "connections.txt"
	err := LoadJSON_Services(FileName)
	if err != nil {
		t.Error("TestLoadJSON_from_file() error: ", err)
	}

	if len(logic.MapServiceURL) == 0 {
		t.Error("TestLoadJSON_from_file() error: len =0")
	}

	//дополнительный файл
	FileName = dir + "settings" + micro.SeparatorFile() + "connections_add.txt"
	err = LoadJSON_Services(FileName)
	if err != nil {
		t.Error("TestLoadJSON_from_file() error: ", err)
	}

	if len(logic.MapServiceURL) == 0 {
		t.Error("TestLoadJSON_from_file() error: len =0")
	}

}
