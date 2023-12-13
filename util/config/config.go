package config

import (
	"log"

	"github.com/go-ini/ini"
)

type zendb struct {
	port string `ini:"port"`
	path string `ini:"path"`
}

var ZendbConfig = new(zendb)

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
	mapTo("zendb", ZendbConfig)
}

// Convert a Map to a struct.
func mapTo(s string, v interface{}) {
	err := cfg.Section(s).MapTo(v)
	if err != nil {
		log.Fatalln(err)
	}
}
