// zendb - client_test.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package client

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	db, err := Connect("127.0.0.1:4780")
	if err != nil {
		fmt.Println(err)
	}
	db.Put("key1key2", "value1value2value3")
}
