package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func Walk(t *tree.Tree, ch1 chan int, start bool) {
	if t != nil {
		Walk(t.Left, ch1, false)
		ch1 <- t.Value
		Walk(t.Right, ch1, false)
		if (start) {
			close(ch1)
		}
	}
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int) 
	go Walk(t1, ch1, true)
	go Walk(t2, ch2, true)
	for {
		x, ok1 := <-ch1 
		y, ok2 := <-ch2 
		switch {
			case ok1 != ok2:
				return false 
			case x != y:
				return false 
			case !ok1 && ok1 == ok2:
				return true
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(tree.New(1), ch1, true)
	go Walk(tree.New(2), ch2, true)
	fmt.Println("tree1")
	for i := range ch1 {
		fmt.Print(i)
	}
	fmt.Println()
	fmt.Println("tree2")
	for j := range ch2 {
		fmt.Print(j)
	}
	fmt.Println()
	fmt.Println("tree1 same as tree2", Same(tree.New(1), tree.New(2)))
	fmt.Println("tree1 same as tree1", Same(tree.New(1), tree.New(1)))
}