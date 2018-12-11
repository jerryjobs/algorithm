package sort

import (
	"fmt"
)

// select sort

// IntsSelectSort sort int array
func IntsSelectSort(in []int) {
	selectSort(IntSlice(in))
}

// Float32SelectSort  float array sort
func Float32SelectSort(in []float32) {
	selectSort(Float32Slice(in))
}

/**
（1）冒泡排序是比较相邻位置的两个数，而选择排序是按顺序比较，找最大值或者最小值；
（2）冒泡排序每一轮比较后，位置不对都需要换位置，选择排序每一轮比较都只需要换一次位置；
（3）冒泡排序是通过数去找位置，选择排序是给定位置去找数；
冒泡排序优缺点：
 		优点:比较简单，空间复杂度较低，是稳定的；
        缺点:时间复杂度太高，效率慢；
选择排序优缺点：
		优点：一轮比较只需要换一次位置；
		缺点：效率慢，不稳定（举个例子5，8，5，2，9   我们知道第一遍选择第一个元素5会和2交换，那么原序列中2个5的相对位置前后顺序就破坏了）。
**/
func selectSort(in Compareble) {

	var times = 1
	l := in.Len()
	for i := 0; i < l; i++ {
		min := i
		for y := i + 1; y < l; y++ {
			if in.More(min, y) {
				min = y
			}
			times++
		}
		if (i + 1) < l {
			in.Swap(i, min)
		}
	}

	if Debug {
		fmt.Printf("for [%d] \n", times)
	}
}
