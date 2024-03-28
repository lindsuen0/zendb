// canodb - log.go
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

func InitLog() {
	err := os.Mkdir("log", 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatalln(err)
	}

	f, err := os.OpenFile("log/canodb.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	Logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
	multiWriter := io.MultiWriter(os.Stdout, f)
	Logger.SetOutput(multiWriter)
}
