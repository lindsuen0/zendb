// canodb - message_test.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package message

import (
	"fmt"
	"testing"
)

func TestParseMess(t *testing.T) {
	s := []byte(":0\n$KEY321\n-VALUE123\n")
	// s := []byte(":0\n$KEY321\n")
	// s := []byte(":0\n$KEY321\n-\n")
	d := parseMess(s, ":", "\n")
	// d := parseMess(s, "$", "\n")
	// d := parseMess(s, "-", "\n")
	fmt.Println(string(d))
}
