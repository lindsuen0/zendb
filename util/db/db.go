package db

import (
	"fmt"
	"log"

	"github.com/lindsuen0/zendb/leveldb"
)

var (
	db  *leveldb.DB
	err error
)

func Setup() {
	db, err = leveldb.OpenFile("data", nil)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("leveldb is on.")
	}
}

// GetData
func GetData(k string) (string, error) {
	v, err := db.Get([]byte(k), nil)
	if err != nil {
		log.Println(err)
	}
	return string(v), err
}

// GetAllData
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

// PutData
func PutData(k string, v string) {
	err := db.Put([]byte(k), []byte(v), nil)
	if err != nil {
		log.Println(err)
	}
}

// DeleteData is a function to delete the value from a key string.
func DeleteData(k string) {
	err := db.Delete([]byte(k), nil)
	if err != nil {
		log.Println(err)
	}
}
