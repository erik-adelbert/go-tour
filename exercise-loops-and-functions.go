// Copyright 2018 Erik Adelbert. All right reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"math"
)

// Sqrt finds the square-root of x, to a precision ε,
// using Newton's method.
func Sqrt(x float64) float64 {
	const ε = 1e-10
	r := float64(1)

	for δ := x; math.Abs(δ) > ε; {
		δ = (r*r - x) / (2 * r)
		r -= δ
	}

	return r
}

// main compares math.Sqrt() and Sqrt().
func main() {
	x := float64(3)
	fmt.Println(Sqrt(x) - math.Sqrt(x))
}
