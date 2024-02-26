// zendb - log.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package log

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

func Setup() {
	err := os.Mkdir("log", 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatalln(err)
	}

	f, err := os.OpenFile("log/logback.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, f)
	Logger = log.New(f, "[ZenDB] ", log.Ldate|log.Lmicroseconds)
	Logger.SetOutput(multiWriter)
}
