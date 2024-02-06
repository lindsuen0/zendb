// zendb - flag.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause License that can be
// found in the LICENSE file.

package flag

import "flag"

type cmdFlag struct {
	statusFlag string
}

// CmdParameter
// ./zendb -s reload, ./zendb -s stop
func CmdParameter(c *cmdFlag) {
	flag.StringVar(&c.statusFlag, "s", "", "The status of ZenDB.")
	flag.Parse()
}
