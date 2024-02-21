// zendb - db.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package db

import (
	"fmt"
	"log"

	"github.com/lindsuen0/zendb/util/config"

	"github.com/lindsuen0/zendb/leveldb"
)

var (
	db  *leveldb.DB
	err error
)

func Setup() {
	dataPath := config.DBConfig.Data
	db, err = leveldb.OpenFile(dataPath, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// GetData is a function to get the value from a key.
func GetData(k string) (string, error) {
	v, err := db.Get([]byte(k), nil)
	if err != nil {
		log.Println(err)
	}
	return string(v), err
}

// GetAllData is a function to get all values in database.
func GetAllData() {
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		fmt.Println("Key: " + string(key))
		fmt.Println("Value: " + string(value))
	}
}

// PutData is a function to add key-value in database.
func PutData(k string, v string) error {
	return db.Put([]byte(k), []byte(v), nil)
}

// DeleteData is a function to delete the value from a key.
func DeleteData(k string) {
	err := db.Delete([]byte(k), nil)
	if err != nil {
		log.Println(err)
	}
}
