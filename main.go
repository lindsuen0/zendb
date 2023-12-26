package main

import (
	"log"
	"net"

	"github.com/lindsuen0/zendb/util/config"
	"github.com/lindsuen0/zendb/util/db"
)

func init() {
	db.Setup()
	config.Setup()
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:4976")
	if err != nil {
		log.Fatalln("Error listening: ", err.Error())
	}
	defer listener.Close()

	log.Println("Server started. Listening on port 4976...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Error accepting connection: ", err.Error())
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// TODO
}
