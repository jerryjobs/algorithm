package sort

import (
	"fmt"
)

// IntsInsertSort insert sort
func IntsInsertSort(in []int) {
	insertSort(IntSlice(in))
}

/**
如果数组中倒置的数量小于数组大小的某个倍数,那么我们说这个数组是部分有序的。
下面是几种典型的部分有序的数组:
	* 数组中每个元素距离它的最终位置都不远;
	* 一个有序的大数组接一个小数组;
	* 数组中只有几个元素的位置不正确。
插入排序对这样的数组很有效,而选择排序则不然。事实上,当倒置的数量很少时是最快的。
*/
func insertSort(in Compareble) {
	l := in.Len()

	if Debug {
		fmt.Printf("len:[%d] \n", l)
	}

	times := 1
	for i := 0; i < l; i++ {
		for j := i; j > 0 && in.More(j-1, j); j-- {
			in.Swap(j-1, j)
			times++
		}
		if Debug {
			fmt.Println(i)
		}
	}
	if Debug {
		fmt.Printf("times :[%d]\n", times)
	}
}
