package sort

import (
	"fmt"
)

// Debug watch for each times
var Debug = false

// BubbleSort bubble sort
func BubbleSort(in []int) []int {

	var size = len(in)
	for i := range in {
		for y := i + 1; y < size; y++ {
			if in[i] > in[y] {
				in[i], in[y] = in[y], in[i]
			}
		}
	}
	return in
}

// Compareble for all compareable.
type Compareble interface {
	// More if from > to then return true else return false.
	More(from, to int) bool
	// Len return lenth of this type.
	Len() int
	// Swap swap for two position.
	Swap(from, to int)
}

// IntSlice for convert int Array.
// (not-a-number values are treated as less than other values).
type IntSlice []int

func (p IntSlice) More(from, to int) bool { return p[from] > p[to] }
func (p IntSlice) Len() int               { return len(p) }
func (p IntSlice) Swap(from, to int)      { p[from], p[to] = p[to], p[from] }

// BubbleSortInts int array sort
func BubbleSortInts(in []int) {
	var a = IntSlice(in)
	sort(a)
}

func IntsIsSorted(in []int) bool {
	return isSorted(IntSlice(in))
}

type Float64Slice []float64

func (p Float64Slice) More(from, to int) bool { return p[from] > p[to] }
func (p Float64Slice) Len() int               { return len(p) }
func (p Float64Slice) Swap(from, to int)      { p[from], p[to] = p[to], p[from] }

// BubbleSortFloat64 sort float64 array
func BubbleSortFloat64(in []float64) {
	sort(Float64Slice(in))
}

func Float64IsSorted(in []float64) bool {
	return isSorted(Float64Slice(in))
}

type Float32Slice []float32

func (p Float32Slice) More(f, t int) bool { return p[f] > p[t] }
func (p Float32Slice) Len() int           { return len(p) }
func (p Float32Slice) Swap(f, t int)      { p[f], p[t] = p[t], p[f] }

// BubbleSortFloat32 sort
func BubbleSortFloat32(in []float32) {
	sort(Float32Slice(in))
}

func Float32IsSorted(in []float32) bool {
	return isSorted(Float32Slice(in))
}

type StringSlice []string

func (s StringSlice) More(f, t int) bool { return s[f] > s[t] }
func (s StringSlice) Len() int           { return len(s) }
func (s StringSlice) Swap(f, t int)      { s[f], s[t] = s[t], s[f] }

// BubbleSortString stor string
func BubbleSortString(in []string) {
	sort(StringSlice(in))
}
func StringIsSorted(in []string) bool {
	return isSorted(StringSlice(in))
}

// every time compare a[x], a[x+1]
// each loop less max position
func sort(c Compareble) {
	l := c.Len()
	var times = 1
	for i := 0; i < l; i++ {
		for y := 0; y < l-i-1; y++ {
			if c.More(y, y+1) {
				c.Swap(y, y+1)
			}
			times++
		}
	}
	if Debug {
		fmt.Printf("for [%d]\n", times)
	}
}

func isSorted(data Compareble) bool {
	for i := 0; i < data.Len()-1; i++ {
		if data.More(i, i+1) {
			return false
		}
	}
	return true
}
