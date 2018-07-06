package web

import (
	"encoding/json"
	"log"
	"mudutv/web/utils"
	"os"
)

var Config JsonFileConfig

type JsonFileConfig struct {
	DatabaseHost string
	DatabasePort string
	DatabaseUser string
	DatabasePass string
	DatabaseName string
}

func init() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	Config = JsonFileConfig{}
	err := decoder.Decode(&Config)

	if err != nil {
		log.Fatalf("加载配置发生致命错误，err:%#v", err)
	}
	utils.InitDB(Config.DatabaseHost, Config.DatabaseUser, Config.DatabasePass,
		Config.DatabasePort, Config.DatabaseName)
}
