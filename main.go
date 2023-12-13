package main

import (
	"github.com/lindsuen0/zendb/util/config"
	"github.com/lindsuen0/zendb/util/db"
)

func init() {
	db.Setup()
	config.Setup()
}

func main() {

}
