package main

import "fmt"
import "time"

const val = 2
const size = 32
const div = size/2

func matMul (done chan [div][div] float64, matX [div][div] float64, matY [div][div] float64)  {
	var matZ [div][div] float64

	for i:= 0; i<div; i++ {
		for j:= 0; j<div; j++ {
			for k:= 0; k<div; k++ {
				matZ[j][i] = matZ[j][i] + ( matX[j][k] * matY[k][i] )
			}
		}
	}

	done <- matZ
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

func matGen() [size][size] float64 {
	var res [size][size] float64

	// filling up the variables
	for i:=0; i<size; i++ {
		for j:=0; j<size; j++ {
			res[i][j] = val
		}
	}

	return res
}

func worker(done chan [div][div] float64, pos int) {
	//var x [div][div] float64
	//var y [div][div] float64

	var matA [size][size] float64 = matGen()
	var matB [size][size] float64 = matGen()

	pos1A := divMat(1, matA)
	pos2A := divMat(2, matA)
	pos3A := divMat(3, matA)
	pos4A := divMat(4, matA)

	pos1B := divMat(1, matB)
	pos2B := divMat(2, matB)
	pos3B := divMat(3, matB)
	pos4B := divMat(4, matB)

	var res [div][div] float64

	mm1 := make(chan [div][div] float64,1)
	mm2 := make(chan [div][div] float64,1)

    if pos == 1 {
    	go matMul(mm1, pos1A, pos1B)
    	go matMul(mm2, pos2A, pos3B)	
    }

    if pos == 2 {
    	go matMul(mm1, pos1A, pos2B)
    	go matMul(mm2, pos2A, pos4B)
    }

    if pos == 3 {
    	go matMul(mm1, pos3A, pos1B)
    	go matMul(mm2, pos4A, pos3B)
    }

    if pos == 4 {
    	go matMul(mm1, pos3A, pos2B)
    	go matMul(mm2, pos4A, pos4B)
    }

	//fmt.Println("run",pos)

    res1 := <- mm1
    res2 := <- mm2
    res = addMat(res1, res2)

	done <- res
}

func main() {

	//parallelization with 4 goroutines
	//TODO: make channel making dynamic
	
	start := time.Now().UnixNano()

	//make channel
	done1 := make(chan [div][div] float64,1)
	done2 := make(chan [div][div] float64,1)
	done3 := make(chan [div][div] float64,1)
	done4 := make(chan [div][div] float64,1)

	go worker(done1, 1)
	go worker(done2, 2)
	go worker(done3, 3)
	go worker(done4, 4)

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