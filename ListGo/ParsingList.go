package main

import (
	"container/list"
	"fmt"
)

func main() {
	neigh := list.New()
	//neigh.PushBack(4)
	//neigh.PushBack(2)

	//begin := neigh.Front()
	//fmt.Println(begin.Value)
	i := 1
	for i < 20 {
		neigh.PushBack(i)
		i++
	}
	// Iterate through list and print its contents.
	for e := neigh.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, ", ")
	}
	fmt.Println()
	fmt.Println(neigh.Len())
	e := neigh.Front()
	for e != nil {
		fmt.Print(e.Value, ", ")
		neigh.Remove(e)
		e = neigh.Front()
	}
	fmt.Println()
	fmt.Println(neigh.Len())

}
