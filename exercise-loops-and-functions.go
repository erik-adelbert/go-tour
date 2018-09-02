// Copyright 2018 Erik Adelbert. All right reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"math"
)

const ε = 1e-10

func Sqrt(x float64) float64 {
	r, δ := x, float64(1)

	for math.Abs(δ) > ε {
		δ = (r*r - x) / (2 * r)
		r -= δ
	}

	return r
}

func main() {
	x := float64(3)
	fmt.Println(Sqrt(x) - math.Sqrt(x))
}
