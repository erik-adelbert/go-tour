// Copyright 2018 Erik Adelbert. All right reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// dfs is a canonical depth first search of tree t.
// it sends all values to the channel ch.
func dfs(t *tree.Tree, ch, quit chan int) {
	if t == nil {
		return
	}

	dfs(t.Left, ch, quit) // smaller values first
	select {
	case ch <- t.Value:
		// nothing more, value sent
	case <-quit:
		return
	}
	dfs(t.Right, ch, quit)
}

// Walk walks the tree t sending all values
// to the channel ch.
func Walk(t *tree.Tree, ch, quit chan int) {
	dfs(t, ch, quit)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	quit := make(chan int)
	defer close(quit)

	go Walk(t1, c1, quit)
	go Walk(t2, c2, quit)

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
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
