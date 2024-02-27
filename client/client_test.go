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

func TestPut(t *testing.T) {
	db, err := Connect("127.0.0.1:4644")
	if err != nil {
		fmt.Println(err)
	}
	db.Put([]byte("key1a"), []byte("value2w"))
}

func TestGet(t *testing.T) {
	db, err := Connect("127.0.0.1:4644")
	if err != nil {
		fmt.Println(err)
	}
	v := db.Get([]byte("key1a"))
	fmt.Println(v)
}
