package config

import (
	"log"

	"github.com/go-ini/ini"
)

type DBConfig struct {
	Port string `ini:"port"`
	Path string `ini:"path"`
}

var ZenDBConfig = new(DBConfig)

var cfg = new(ini.File)

func Setup() {
	var configPath = "config/config.ini"
	var err error

	cfg, err = ini.Load(configPath)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("%s is parsed successfully.", configPath)
	}
	mapTo("zendb", ZenDBConfig)
}

// Convert a Map to a struct.
func mapTo(s string, v interface{}) {
	err := cfg.Section(s).MapTo(v)
	if err != nil {
		log.Fatalln(err)
	}
}
