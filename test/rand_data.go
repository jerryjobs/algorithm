package test

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"jerrysd.cn/algorithm/array/sort"
	"jerrysd.cn/algorithm/test"
)

// TestRand rand values for sort
func Test_Rand(t *testing.T) {

	var a = make([]int, 9999999)
	for i := 0; i < 9999999; i++ {
		a[i] = i
	}

	for i := 0; i < 9999999; i++ {
		v := rand.Intn(9999999)
		next := rand.Intn(9999999)
		temp := a[v]
		a[v] = a[next]
		a[next] = temp
	}

	//os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	f, err := os.OpenFile("input.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		t.Fatal(err)
	} else {
		for _, v := range a {
			//f.WriteString(fmt.Sprintf("%d\n", v))
			fmt.Fprintln(f, v)
		}
		f.Close()
		fmt.Println("ok")
	}
}

//ReadArrayFromFile read file and read data
func ReadArrayFromFile(fileName string) []int {

	f, err := os.Open(fileName)

	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(f)

	item, err := reader.ReadString('\n')
	index := 1

	var rs = make([]int, 0)

	for err != io.EOF {
		index++
		//fmt.Printf("%d \t- %s", index, item)
		if i, err := strconv.Atoi(strings.Trim(item, "\n")); err == nil {
			//fmt.Printf("%d \t- %d \n", index, i)
			rs = append(rs, i)
		} else {
			log.Fatal(err)
			break
		}

		item, err = reader.ReadString('\n')
	}

	if err != nil && err != io.EOF {
		panic(err)
	}
	return rs
}

// Test_insertionSort
func Test_InsertionSort(t *testing.T) {
	arr := test.ReadArrayFromFile("resource/input2.txt")

	tu := time.Now()
	sort.Debug = true
	sort.IntsInsertSort(arr)
	rs := time.Now().Sub(tu) / time.Microsecond
	fmt.Println("--", rs)
	if sort.IntsIsSorted(arr) {
		fmt.Println("ok")
	} else {
		fmt.Println("false")
	}
}
