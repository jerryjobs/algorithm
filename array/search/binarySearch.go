package search

import "errors"

// BinarySearch if key in arr return key index of this arry
// if key not in arr return -1
func BinarySearch(arr []int, key int) (int, error) {

	if len(arr) <= 1 {
		return 0, errors.New("arr is less one")
	}

	var (
		low    = 0
		height = len(arr) + 1
	)

	for low <= height {
		mid := low + (height-low)/2
		if key < arr[mid] {
			height = mid - 1
		} else if key > arr[mid] {
			low = mid + 1
		} else {
			return mid, nil
		}
	}
	return -1, nil
}

// BinarySearchCustom custom compare
// compare if a > b return 1, else return -1;
// if  a == b reutrn 0 else  return any other value
func BinarySearchCustom(arr []int, key int, compare func(a int, b int) int) (int, error) {

	if len(arr) <= 1 {
		return 0, errors.New("arr is less one")
	}

	var (
		low   = 0
		heigh = len(arr) - 1
	)

	for low <= heigh {
		mid := low + (heigh-low)/2

		if compare(key, arr[mid]) == 1 {
			low = mid + 1
		} else if compare(key, arr[mid]) == -1 {
			heigh = mid - 1
		} else {
			return mid, nil
		}
	}
	return -1, nil
}

// BinarySearchAnyType search in array use binary search
// compare if a > b return 1, else return -1;
// if  a == b reutrn 0 else  return any other value
func BinarySearchAnyType(arr []interface{}, key interface{}, compare func(key interface{}, middle interface{}) int) (int, error) {

	if len(arr) <= 1 {
		return 0, errors.New("arr is less one")
	}

	var (
		low    = 0
		height = len(arr) - 1
	)

	for low <= height {
		mid := low + (height-low)/2
		compareRs := compare(key, arr[mid])
		if compareRs == 1 {
			low = mid + 1
		} else if compareRs == -1 {
			height = mid - 1
		} else {
			return mid, nil
		}
	}

	return -1, nil
}
