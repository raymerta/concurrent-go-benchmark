package main

import "fmt"
import "time"

const hashSize = 5

func genInput(start int, stop int) [hashSize] int {
	var arr [hashSize] int

	for i:= 0; i < hashSize; i++ {
		arr[i] = i + start
	}

	return arr
}

func main() {

	start := time.Now().UnixNano()

	var arr [hashSize] int = genInput(10, 100)
	fmt.Println(arr)

	stop := time.Now().UnixNano()
	fmt.Println((stop - start))
}