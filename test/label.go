// errorcheck

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Verify that erroneous labels are caught by the compiler.
// This set is caught by pass 1.
// Does not compile.

package main

var x int

func f() {
L1:
	for {
	}
L2:
	select {}
L3:
	switch {
	}
L4:
	if true {
	}
L5:
	f()
L6: // GCCGO_ERROR "previous"
	f()
L6: // ERROR "label .*L6.* already defined"
	f()
	if x == 20 {
		goto L6
	}

L7:
	for {
		break L7
	}

L8:
	for {
		if x == 21 {
			continue L8
		}
	}

L9:
	switch {
	case true:
		break L9
	defalt:
	}

L10:
	select {
	default:
		break L10
	}

	goto L10

	goto go2 // ERROR "label go2 not defined"
}
