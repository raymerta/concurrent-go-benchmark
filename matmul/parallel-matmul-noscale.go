package main

import "fmt"
import "time"

const val = 2
const size = 4
const div = size/2

func matMul (matX [div][div] float64, matY [div][div] float64) [div][div] float64 {
	var matZ [div][div] float64

	for i:= 0; i<div; i++ {
		for j:= 0; j<div; j++ {
			for k:= 0; k<div; k++ {
				matZ[j][i] = matZ[j][i] + ( matX[j][k] * matY[k][i] )
			}
		}
	}

	return matZ
}

func addMat(matX [div][div] float64, matY [div][div] float64) [div][div] float64 {
	var matZ [div][div] float64

	for i:= 0; i < div; i++ {
		for j:=0; j < div; j++ {
			matZ[i][j] = matX[i][j] + matY[i][j]
		}
	}

	return matZ
}


func divMat(pos int, mat [size][size] float64) [div][div] float64 {

	var res [div][div] float64

	for i:=0; i < div; i++ {
		for j:=0; j < div; j++ {
			if pos == 1 {
				res[i][j] = mat[i][j]
			}

			if (pos == 2) {
				res[i][j] = mat[i][j + div]
		   	}

			if (pos == 3) {
			    res[i][j] = mat[i + div][j]
			}

			if (pos == 4) {
			    res[i][j] = mat[i + div][j + div]
			}

		}
	}

	return res
}

func assMat(pos int, mat [div][div] float64, res [size][size] float64) [size][size] float64 {

	for i:=0; i < div; i++ {
		for j:=0; j < div; j++ {
			if pos == 1 {
				res[i][j] = mat[i][j]
				//fmt.Println(res[i][j])
			}

			if pos == 2 {
				res[i][j + div] = mat[i][j]
				//fmt.Println(res[i][j + div])
		   	}

			if pos == 3 {
				res[i + div][j] = mat[i][j]
				//fmt.Println(res[i + div][j])
			}

			if pos == 4 {
				res[i + div][j + div] = mat[i][j]
				//fmt.Println(res[i + div][j + div])
			}

		}
	}

	//fmt.Println(res)
	return res
}

func worker(done chan [div][div] float64, pos int, matA [size][size] float64, matB [size][size] float64) {
	//var x [div][div] float64
	//var y [div][div] float64

	pos1A := divMat(1, matA)
	pos2A := divMat(2, matA)
	pos3A := divMat(3, matA)
	pos4A := divMat(4, matA)

	pos1B := divMat(1, matB)
	pos2B := divMat(2, matB)
	pos3B := divMat(3, matB)
	pos4B := divMat(4, matB)

	var res [div][div] float64

    if pos == 1 {
    	res = addMat(matMul(pos1A, pos1B), matMul(pos2A, pos3B))
    }

    if pos == 2 {
    	res = addMat(matMul(pos1A, pos2B), matMul(pos2A, pos4B))
    }

    if pos == 3 {
    	res = addMat(matMul(pos3A, pos1B), matMul(pos4A, pos3B))
    }

    if pos == 4 {
    	res = addMat(matMul(pos3A, pos2B), matMul(pos4A, pos4B))
    }

	//fmt.Println("run",pos)

	done <- res
}

func main() {

	//declare matrix container
	var a [size][size] float64
	var b [size][size] float64

	// filling up the variables
	for i:=0; i<size; i++ {
		for j:=0; j<size; j++ {
			a[i][j] = val
			b[i][j] = val
		}
	}

	//parallelization with 4 goroutines
	//TODO: make channel making dynamic
	
	start := time.Now().UnixNano()

	//make channel
	done1 := make(chan [div][div] float64,1)
	done2 := make(chan [div][div] float64,1)
	done3 := make(chan [div][div] float64,1)
	done4 := make(chan [div][div] float64,1)

	go worker(done1, 1, a, b)
	go worker(done2, 2, a, b)
	go worker(done3, 3, a, b)
	go worker(done4, 4, a, b)

	res1 := <-done1
	res2 := <-done2
	res3 := <-done3
	res4 := <-done4

	var c [size][size] float64

	c = assMat(1, res1, c)
	c = assMat(2, res2, c)
	c = assMat(3, res3, c)
	c = assMat(4, res4, c)

	stop := time.Now().UnixNano()
	fmt.Println((stop - start))
	fmt.Println(c)
}