package main

import (
	"fmt"
	"math/rand"
)

func makeRandomNumsNegative(max, size int) []bool {
	double := 2 * max	
	rand.Seed(int64(double))
	var boolArray []bool

	for i := 0; i < size; i++ {
		tmp := int(rand.Int63n(int64(double)))
		if tmp > max {
			boolArray = append(boolArray, false)
		} else {
			boolArray = append(boolArray, true)
		}
	}

	return boolArray
}

func fillArray(max, size int) []int {
	rand.Seed(int64(max))
	var arr []int
	negative := makeRandomNumsNegative(max, size)

	for i := 0; i < size; i++ {
		tmp := int(rand.Int63n(int64(max)))
		if negative[i] {
			arr = append(arr, -tmp)
		} else {
			arr = append(arr, tmp)
		}
	}

	return arr
}

func maxSubarray(a []int) (int, int, int) {
	if a == nil {
		return -1, -1, 0 
	}

	maxRes, maxHere := a[0], a[0]
	size := len(a)
	maxBegin, maxEnd, tmp := 0, 0, 0

	for i := 1; i < size; i++ {
		maxHere += a[i]
		if maxRes < maxHere {
			maxBegin = tmp
			maxEnd = i
			maxRes = maxHere
		}
		if maxHere < 0 {
			maxHere = 0
			tmp = i + 1
		}
	}

	return maxBegin, maxEnd, maxRes
}

func main() {
	var size int
	fmt.Print("Enter size of array: ")
	fmt.Scan(&size)

	var max int
	fmt.Print("Enter maximum number which may be in array: ")
	fmt.Scan(&max)

	arr := fillArray(max, size)

	fmt.Println(arr)

	subarrayBegin, subArrayEnd, sum := maxSubarray(arr)

	if subarrayBegin == -1 && subArrayEnd == -1 {
		fmt.Println("Slice is empty.")
	} else {
		subarray := arr[subarrayBegin:(subArrayEnd + 1)]
		fmt.Println(subarray)
		fmt.Printf("Maximum sum in this subarray is %d\n", sum)
	}
}
