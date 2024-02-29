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
	}
	fmt.Println(string(v))
}
