// Copyright 2018 Erik Adelbert. All right reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	for _, w := range strings.Fields(s) {
		m[w]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
