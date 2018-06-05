package main

import "fmt"

func main() {

	const val = 2
	const size = 3

	var a [size][size] int
	var b [size][size] int
	var c [size][size] int

	// filling up the variables
	for i:=0; i<size; i++ {
		for j:=0; j<size; j++ {
			a[i][j] = val
			b[i][j] = val
		}
	}

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

	fmt.Println("c: ", c)
}
