package main

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	src := []int{9, 2, 3, 4, 6, 8, 10, 23, 4, 43, 5}
	BubbleSort(src)
	fmt.Println(src)
}

func TestMergeSort(t *testing.T) {
	src := []int{9, 2, 3, 4, 6, 8, 10, 23, 4, 43, 5}
	MergeSortInPlace(src, 0, len(src))
	fmt.Println(src)
}

func TestMergeSortGoroutine(t *testing.T) {
	src := []int{9, 2, 3, 4, 6, 8, 10, 23, 4, 43, 5}
	fmt.Println(src)
	exit := make(chan struct{}, 0)
	go MergeSortInPlaceGoroutine(src, 0, len(src)-1, exit)
	<-exit
	fmt.Println(src)
}
