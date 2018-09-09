// Copyright 2018 Erik Adelbert. All right reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walker search the tree t sending all values to channel ch.
func Walker(t *tree.Tree, quit chan int) <-chan int {
	ch := make(chan int)
	go func() {
		walk(t, ch, quit)
		close(ch)
	}()
	return ch
}

// walk is a canonical depth first search of tree t.
// It sends all values to the channel ch.
func walk(t *tree.Tree, ch, quit chan int) {
	if t == nil {
		return
	}
	
	walk(t.Left, ch, quit) // smaller values first
	select {
	case ch <- t.Value:
		// nothing more, value sent
	case <-quit:
		return
	}
	walk(t.Right, ch, quit)
}

// Same determines whether the trees t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	quit := make(chan int)
	defer close(quit)

	c1 := Walker(t1, quit)
	c2 := Walker(t2, quit)

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		if !ok1 || !ok2 {
			return ok1 == ok2
		}

		if v1 != v2 {
			return false
		}
	}
	
	return false
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
