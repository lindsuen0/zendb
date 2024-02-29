// canodb - stream_test.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package stream

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

func TestParseMess(t *testing.T) {
	s := []byte(":0\n$KEY321\n-VALUE123\n")
	// s := []byte(":0\n$KEY321\n")
	// s := []byte(":0\n$KEY321\n-\n")
	d := parseMess(s, ":", "\n")
	// d := parseMess(s, "$", "\n")
	// d := parseMess(s, "-", "\n")
	fmt.Println(string(d))
}
