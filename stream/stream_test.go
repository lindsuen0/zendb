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
	stream1 := CreatePutMess([]byte("key1"), []byte("value1"))
	operatorStream1 := append(append(stream1.Operator.StartTag, stream1.Operator.OperatorContent...), stream1.Operator.EndTag...)
	keyStream1 := append(append(stream1.Key.StartTag, stream1.Key.KeyContent...), stream1.Key.EndTag...)
	valueStream1 := append(append(stream1.Value.StartTag, stream1.Value.ValueContent...), stream1.Value.EndTag...)
	fmt.Printf("%q", string(append(append(operatorStream1, keyStream1...), valueStream1...)))
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
