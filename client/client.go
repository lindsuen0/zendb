// canodb - client.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package client

import (
	"fmt"
	s "github.com/lindsuen/canodb/stream"
	"net"
	"time"
)

var (
	READTIMEOUT = 3 * time.Second
)

type Driver struct {
	Conn net.Conn
}

func Connect(tcpAddress string) (*Driver, error) {
	conn, err := net.Dial("tcp", tcpAddress)
	return &Driver{conn}, err
}

func (n *Driver) Put(key []byte, value []byte) error {
	p := s.CreatePutMess(key, value)
	operatorString := mergeByteSlice(p.Operator.StartTag, p.Operator.OperatorContent, p.Operator.EndTag)
	keyString := mergeByteSlice(p.Key.StartTag, p.Key.KeyContent, p.Key.EndTag)
	valueString := mergeByteSlice(p.Value.StartTag, p.Value.ValueContent, p.Value.EndTag)

	_, err := n.Conn.Write(append(append(operatorString, keyString...), valueString...))
	return err
}

func (n *Driver) Delete(key []byte) error {
	p := s.CreateDelMess(key)
	operatorString := mergeByteSlice(p.Operator.StartTag, p.Operator.OperatorContent, p.Operator.EndTag)
	keyString := mergeByteSlice(p.Key.StartTag, p.Key.KeyContent, p.Key.EndTag)

	_, err := n.Conn.Write(append(operatorString, keyString...))
	return err
}

func (n *Driver) Get(key []byte) ([]byte, error) {
	var err error
	p := s.CreateGetMess(key)
	operatorString := mergeByteSlice(p.Operator.StartTag, p.Operator.OperatorContent, p.Operator.EndTag)
	keyString := mergeByteSlice(p.Key.StartTag, p.Key.KeyContent, p.Key.EndTag)

	_, err = n.Conn.Write(append(operatorString, keyString...))
	if err != nil {
		fmt.Println(err)
	}

	_ = n.Conn.SetReadDeadline(time.Now().Add(READTIMEOUT))
	buf := [256]byte{}
	b, _ := n.Conn.Read(buf[:])
	return buf[:b], err
}

func mergeByteSlice(startTag []byte, content []byte, endTag []byte) []byte {
	return append(append(startTag, content...), endTag...)
}
