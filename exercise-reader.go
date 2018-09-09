// Copyright 2018 Erik Adelbert. All right reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// MyReader.Read returns an infinite stream of 'A'.
// It does so len(b) bytes at a time.
func (r MyReader) Read(b []byte) (int, error) {
	
	//The core loop fills b with as many 'A' as possible.
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
