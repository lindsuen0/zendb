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

type dbConfig struct {
	Port string `ini:"port"`
	Data string `ini:"data"`
}

var DBConfig = new(dbConfig)

var cfg = new(ini.File)

func Setup() {
	var configPath = "config/config.ini"
	var err error

	cfg, err = ini.Load(configPath)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("The profile %s is parsed successfully.\n", configPath)
	}
	mapTo("db", DBConfig)
}

// Convert a Map to a struct.
func mapTo(s string, v interface{}) {
	err := cfg.Section(s).MapTo(v)
	if err != nil {
		log.Fatalln(err)
	}
}
