// Copyright 2018 Erik Adelbert. All right reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rotn(n, c byte) byte {
	var a, z byte

	switch {
	case 'a' <= c && c <= 'z':
		a, z = 'a', 'z'
	case 'A' <= c && c <= 'Z':
		a, z = 'A', 'Z'
	default:
		return c
	}
	return a + (c-a+n)%(1+z-a)
}

func (m rot13Reader) Read(b []byte) (int, error) {
	n, err := m.r.Read(b)

	if err == nil {
		for i, v := range b {
			b[i] = rotn(13, v)
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
