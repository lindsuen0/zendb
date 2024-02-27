// canodb - client.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package client

import (
	"fmt"
	"net"
	"os"

	s "github.com/lindsuen0/canodb/stream"
)

type Driver struct {
	Connection net.Conn
	Hostname   string
}

func Connect(tcpAddress string) (*Driver, error) {
	conn, err := net.Dial("tcp", tcpAddress)
	hostname, _ := os.Hostname()
	// defer conn.Close()

	// inputReader := bufio.NewReader(os.Stdin)
	// for {
	// 	input, _ := inputReader.ReadString('\n')
	// 	inputInfo := strings.Trim(input, "\r\n")
	// 	if strings.ToUpper(inputInfo) == "Q" {
	// 		return
	// 	}
	// 	_, err := conn.Write([]byte(inputInfo))
	// 	if err != nil {
	// 		return
	// 	}
	// 	buf := [512]byte{}
	// 	n, err := conn.Read(buf[:])
	// 	if err != nil {
	// 		log.Println("Recvied failed, err: ", err)
	// 		return
	// 	}
	// 	log.Println(string(buf[:n]))
	// }
	return &Driver{conn, hostname}, err
}

func (n *Driver) Put(key []byte, value []byte) {
	p := s.GeneratePutStream(key, value)
	operatorString := mergeString(p.Operator.StartTag, p.Operator.OperatorContent, p.Operator.EndTag)
	keyString := mergeString(p.Key.StartTag, p.Key.KeyContent, p.Key.EndTag)
	valueString := mergeString(p.Value.StartTag, p.Value.ValueContent, p.Value.EndTag)

	_, err := n.Connection.Write(append(append(operatorString, keyString...), valueString...))
	if err != nil {
		fmt.Println(err)
	}
}

func (n *Driver) Delete(key []byte) {
	p := s.GenerateGetStream(key)
	operatorString := mergeString(p.Operator.StartTag, p.Operator.OperatorContent, p.Operator.EndTag)
	keyString := mergeString(p.Key.StartTag, p.Key.KeyContent, p.Key.EndTag)

	_, err := n.Connection.Write(append(append(operatorString, keyString...)))
	if err != nil {
		fmt.Println(err)
	}
}

func (n *Driver) Get(key []byte) string {
	p := s.GenerateGetStream(key)
	operatorString := mergeString(p.Operator.StartTag, p.Operator.OperatorContent, p.Operator.EndTag)
	keyString := mergeString(p.Key.StartTag, p.Key.KeyContent, p.Key.EndTag)

	_, err := n.Connection.Write(append(append(operatorString, keyString...)))
	if err != nil {
		fmt.Println(err)
	}

	buf := [512]byte{}
	b, _ := n.Connection.Read(buf[:])
	return string(buf[:b])
}

func mergeString(startTag []byte, content []byte, endTag []byte) []byte {
	return append(append(startTag, content...), endTag...)
}
