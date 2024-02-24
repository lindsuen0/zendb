// zendb - client_test.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package client

import "testing"

func TestConnect(t *testing.T) {
	db := Connect("127.0.0.1:4780")
	db.Put("key1ssssssss", "value2sssssss1")
}
