// zendb - log.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package log

import (
	"github.com/lindsuen0/zendb/util/config"
	"log"
	"os"
	"time"
)

var timeStr = time.Now().Format("2006-01-02 15:04:05")

func Setup() {
	err := os.Mkdir(config.DBConfig.Log, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatalln(err)
	} else {
		Println("The directory of log has been created.")
	}
}

func Println(v string) {
	err := writeFile(config.DBConfig.Log+"/logback.log", timeStr+" [ZenDB] "+v, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("[ZenDB] " + v)
}

// func Printf(format string, v ...any) {
// 	err := writeFile(config.DBConfig.Log+"/logback.log", fmt.Sprintf(timeStr+" [ZenDB] %s", v), 0666)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	log.Printf("[ZenDB] "+format, v)
// }

func writeFile(name string, data string, perm os.FileMode) error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND|os.O_CREATE, perm)
	if err != nil {
		return err
	}
	_, err = f.WriteString(data + "\n")
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}
