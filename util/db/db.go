// canodb - db.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package db

import (
	"log"

	"github.com/lindsuen0/canodb/leveldb"
	c "github.com/lindsuen0/canodb/util/config"
	l "github.com/lindsuen0/canodb/util/log"
)

var (
	db  *leveldb.DB
	err error
)

func Setup() {
	dataPath := c.DBConfig.Data
	db, err = leveldb.OpenFile(dataPath, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func PutData(k []byte, v []byte) error {
	l.Logger.Println("key: \"" + string(k) + "\", value: \"" + string(v) + "\"")
	return db.Put(k, v, nil)
}

func DeleteData(k []byte) {
	err := db.Delete(k, nil)
	if err != nil {
		l.Logger.Println(err)
	}
}

func GetData(k []byte) ([]byte, error) {
	v, err := db.Get(k, nil)
	if err != nil {
		l.Logger.Println(err)
	}
	return v, err
}
