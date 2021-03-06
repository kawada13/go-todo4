package config

import (
	"log"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

var Config ConfigList

// main関数が呼ばれる前に実行される
func init() {
	LoadConfig()
}

func LoadConfig() {
	// Loadするタイミングで、logファイルが存在しなければ、作成する
	cfg, err := ini.Load("config.ini")

	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}

}
