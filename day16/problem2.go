package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func generateCoefficients(k int, N int) []int {
	// generation 1: 1 1  1  1  1 ... (N times)
	//            2: 1 2  3  4  5 ...
	// generation 3: 1 3  6 10 15 ...
	// generation 4: 1 4 10 20 35 ...
	//            5: 1 5 15 35 70

	c := make([]int, N)
	for i := 0; i < N; i++ {
		c[i] = 1
	}
	for k > 1 {
		for i := 1; i < N; i++ {
			c[i] += c[i-1]
			c[i] %= 10
		}
		k--
	}
	return c
}

func calculateOffset(data []int) int {
	offset := 0
	for i := 0; i < 7; i++ {
		offset = 10*offset + data[i]
	}
	return offset
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)

	data := make([]int, 0)
	for s.Scan() {
		digit, err := strconv.Atoi(s.Text())
		if err != nil {
			continue
		}
		data = append(data, digit)
	}

	// calculate coefficients using generating function.
	// This relies on the fact that the FFT operation is multiplication with a matrix
	//
	//  A = [a b]
	//      [0,1]
	//
	// where the lower right quadrant is an upper-triangular matrix consisting
	// of 1s on and above the diagonal. Doing the FFT operation 100 times
	// preserves the structure of A and turns the lower right quadrant into an
	// upper triangular matrix with the following structure
	//
	//  I = [ x y z u ... ]
	//      [ 0 x y z ... ]
	//      [ 0 0 x y ... ]
	//
	// The coefficients x, y, z, u, ... can be calculated explicitly, this is
	// done by the function generateCoefficients.

	seqlen := len(data) * 10000
	offset := calculateOffset(data)
	nonzero := seqlen + 1 - offset
	c := generateCoefficients(100, nonzero)

	for i := 0; i < 8; i++ {
		sum := 0
		for j := offset + i; j < seqlen; j++ {
			el := data[j%len(data)]
			sum += c[j-offset-i] * el
			sum %= 10
		}
		fmt.Print(sum)
	}
	fmt.Println()
}
