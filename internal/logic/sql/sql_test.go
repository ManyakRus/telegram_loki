package sql

import (
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/micro"
	"github.com/ManyakRus/starter/postgres_pgx"
	"github.com/ManyakRus/telegram_loki/internal/config"
	"testing"
)

func TestRunSQL1(t *testing.T) {
	config_main.LoadEnv()
	config.FillSettings()

	postgres_pgx.Connect()
	defer postgres_pgx.CloseConnection()

	dir := micro.ProgramDir()
	Filename := dir + "test" + micro.SeparatorFile() + "test_error" + micro.SeparatorFile() + "test.sql"
	err := RunSQL1(Filename)
	if err == nil {
		t.Error("TestRunSQL1() error: ", err)
	} else {
		t.Log("TestRunSQL1() result: ", err)
	}
}
