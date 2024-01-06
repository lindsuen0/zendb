// zendb - config.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package config

import (
	"log"

	"github.com/go-ini/ini"
)

type DBConfig struct {
	ZenDBPort string `ini:"port"`
	ZenDBData string `ini:"data"`
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
