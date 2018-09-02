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

func rot13(c byte) byte {
	σ := map[byte]byte { // ROT13 permutation is not so big
		'A': 'N', 'B': 'O', 'C': 'P', 'D': 'Q',
		'E': 'R', 'F': 'S', 'G': 'T', 'H': 'U',
		'I': 'V', 'J': 'W', 'K': 'X', 'L': 'Y',
		'M': 'Z', 'N': 'A', 'O': 'B', 'P': 'C',
		'Q': 'D', 'R': 'E', 'S': 'F', 'T': 'G',
		'U': 'H', 'V': 'I', 'W': 'J', 'X': 'K',
		'Y': 'L', 'Z': 'M', 'a': 'n', 'b': 'o',
		'c': 'p', 'd': 'q', 'e': 'r', 'f': 's',
		'g': 't', 'h': 'u', 'i': 'v', 'j': 'w',
		'k': 'x', 'l': 'y', 'm': 'z', 'n': 'a',
		'o': 'b', 'p': 'c', 'q': 'd', 'r': 'e',
		's': 'f', 't': 'g', 'u': 'h', 'v': 'i',
		'w': 'j', 'x': 'k', 'y': 'l', 'z': 'm',
	}

	if r, ok := σ[c]; ok {
		return r
	}

	return c
}

func (m rot13Reader) Read(b []byte) (n int, err error) {

	if n, err = m.r.Read(b); err == nil {
		for i, v := range b {
			b[i] = rot13(v)
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
