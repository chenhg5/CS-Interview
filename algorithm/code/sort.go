package main

import "fmt"

// 排序算法集合
// --------------------------
// > 冒泡排序
// > 选择排序
// > 插入排序
// > 希尔排序：插入的升级
// > 归并排序
// > 快速排序：冒泡的升级
// > 堆排序
// > 计数排序
// > 桶排序：计数的升级
// > 基数排序
// > 外排序

// O(1) > O(n) > O(logn) > O(n^2) > O(2^n)
// 常数    线性    对数      多项式     指数

// 9, 2, 3, 4, 6, 8, 10, 23, 4, 43, 5
// 2, 3, 4, 6, 8, 9, 10, 23, 4, 5,             43
// 2, 3, 4, 6, 8, 9, 10, 4, 5              23, 43

func BubbleSort(src []int) {
	fmt.Println("bubble sort")

	bigOne := 0

	fmt.Println(src)

	for j := 0; j < len(src); j++ {
		for i := 0; i < len(src)-j; i++ {
			if src[bigOne] < src[i] {
				bigOne = i
			}
			if i == len(src)-j-1 {
				src[bigOne], src[len(src)-j-1] = src[len(src)-j-1], src[bigOne]
				fmt.Println(src)
			}
		}
	}

	// n + (n-1) + (n-2) + ... + 1 = (n + 1) * n / 2 = O(n^2)
}

// 合并 [l,r] 两部分数据，mid 左半部分的终点，mid + 1 是右半部分的起点
func merge(arr []int, l int, mid int, r int) {
	// 因为需要直接修改 arr 数据，这里首先复制 [l,r] 的数据到新的数组中，用于赋值操作
	temp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = arr[i]
	}

	// 指向两部分起点
	left := l
	right := mid + 1

	for i := l; i <= r; i++ {
		// 左边的点超过中点，说明只剩右边的数据
		if left > mid {
			arr[i] = temp[right-l]
			right++
			// 右边的数据超过终点，说明只剩左边的数据
		} else if right > r {
			arr[i] = temp[left-l]
			left++
			// 左边的数据大于右边的数据，选小的数字
		} else if temp[left-l] > temp[right-l] {
			arr[i] = temp[right-l]
			right++
		} else {
			arr[i] = temp[left-l]
			left++
		}
	}
}

func MergeSortInPlace(arr []int, l int, r int) {
	if l >= r {
		return
	}

	// 递归向下
	mid := (r + l) / 2
	MergeSortInPlace(arr, l, mid)
	MergeSortInPlace(arr, mid+1, r)
	// 归并向上
	merge(arr, l, mid, r)
	fmt.Println("l", l, "r", r, "arr", arr)
}

func MergeSortInPlaceGoroutine(arr []int, l int, r int, exit chan struct{}) {
	if l >= r {
		exit <- struct{}{}
		return
	}

	// 递归向下
	mid := (r + l) / 2
	var a1 = make(chan struct{}, 0)
	var a2 = make(chan struct{}, 0)
	go MergeSortInPlaceGoroutine(arr, l, mid, a1)
	go MergeSortInPlaceGoroutine(arr, mid+1, r, a2)
	<-a1
	<-a2
	// 归并向上
	merge(arr, l, mid, r)
	fmt.Println("l", l, "r", r, "arr", arr)
	exit <- struct{}{}
}
