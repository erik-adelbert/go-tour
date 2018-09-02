// Copyright 2018 Erik Adelbert. All right reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import "golang.org/x/tour/pic"

// Pic computes and returns a dx*dy bitmap.
func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)         // make space for dy lines
	for y := range p {
		p[y] = make([]uint8, dx) // make dx columns 
		for x := range p[y] {    // fill the line with (blue) color levels 
			p[y][x] = uint8(x ^ y - y ^ x)
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}
