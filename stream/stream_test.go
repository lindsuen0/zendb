// zendb - stream_test.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package stream

import (
	"fmt"
	"testing"
)

func TestGeneratePutStream(t *testing.T) {
	stream1 := GeneratePutStream("key1", "value1")
	operatorStream1 := stream1.Operator.StartTag + stream1.Operator.OperatorContent + stream1.Operator.EndTag
	keyStream1 := stream1.Key.StartTag + stream1.Key.KeyContent + stream1.Key.EndTag
	valueStream1 := stream1.Value.StartTag + stream1.Value.ValueContent + stream1.Value.EndTag
	fmt.Println(operatorStream1 + keyStream1 + valueStream1)
}

func TestGenerateDeleteStream(t *testing.T) {
	GenerateDeleteStream("key1")
}

func TestParseStruct(t *testing.T) {
	s := ":0\n$KEY321\n-VALUE123\n"
	// d := parseStruct(s, ":", "\n")
	// d := parseStruct(s, "$", "\n")
	d := parseStruct(s, "-", "\n")
	fmt.Println(d)
}
