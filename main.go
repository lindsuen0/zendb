package main

import (
	"bufio"
	"fmt"
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
	listener, err := net.Listen("tcp", "127.0.0.1:"+config.ZenDBConfig.ZenDBPort)
	if err != nil {
		log.Fatalln("Error listening: ", err.Error())
	}
	defer listener.Close()

	log.Println("Server started. Listening on port " + config.ZenDBConfig.ZenDBPort + "...")

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
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err: ", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("have recived msg from client: ", recvStr)
		conn.Write([]byte(recvStr))
	}
}
