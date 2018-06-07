package main

import "fmt"
import "time"

func main() {

	const val = 2
	const size = 32

	//declare matrix container
	var a [size][size] float64
	var b [size][size] float64
	var c [size][size] float64

	// filling up the variables
	for i:=0; i<size; i++ {
		for j:=0; j<size; j++ {
			a[i][j] = val
			b[i][j] = val
		}
	}

	//start counter
	start := time.Now().UnixNano()

	//test if the content is correct
	//fmt.Println("a: ", a)
	//fmt.Println("b: ", b)

	for i:= 0; i<size; i++ {
		for j:= 0; j<size; j++ {
			for k:= 0; k<size; k++ {
				c[j][i] = c[j][i] + ( a[j][k] * b[k][i] )
			}
		}
	}

	stop := time.Now().UnixNano()
	fmt.Println((stop - start))
	fmt.Println("c: ", c)
}
