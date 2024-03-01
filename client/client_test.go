// canodb - client_test.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package client

import (
	"fmt"
	"testing"
)

func TestCreatePutMess(t *testing.T) {
	message := CreatePutMess([]byte("key1"), []byte("value1"))
	operatorMessage := append(append(message.Operator.StartTag, message.Operator.OperatorContent...), message.Operator.EndTag...)
	keyMessage := append(append(message.Key.StartTag, message.Key.KeyContent...), message.Key.EndTag...)
	valueMessage := append(append(message.Value.StartTag, message.Value.ValueContent...), message.Value.EndTag...)
	fmt.Printf("%q", string(append(append(operatorMessage, keyMessage...), valueMessage...)))
}

var (
	db  *Driver
	err error
)

func TestPut(t *testing.T) {
	db, err = Connect("127.0.0.1:4644")
	if err != nil {
		fmt.Println(err)
	}

	err = db.Put([]byte("key1a"), []byte("ddvalue2w"))
	if err != nil {
		fmt.Println(err)
	}
}

func TestDelete(t *testing.T) {
	db, err = Connect("127.0.0.1:4644")
	if err != nil {
		fmt.Println(err)
	}

	err = db.Delete([]byte("key1a"))
	if err != nil {
		fmt.Println(err)
	}
}

func TestGet(t *testing.T) {
	db, err = Connect("127.0.0.1:4644")
	if err != nil {
		fmt.Println(err)
	}

	var v []byte
	v, err = db.Get([]byte("key1a"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(v))
	}
}
