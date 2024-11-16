package bigo

import "fmt"

// A * B, A and B are matrics
func multiplyMatrices(a, b [][]int) [][]int {
	rowsA, colsA := len(a), len(a[0])
	rowsB, colsB := len(b), len(b[0])

	if colsA != rowsB {
		panic("Matrices cannot be multiplied")
	}

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	for i := 0; i < rowsA; i++ { // first loop
		for j := 0; j < colsB; j++ { // second loop
			for k := 0; k < colsA; k++ { // third loop
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return result
}

func mainn3() {

	matrixA := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	matrixB := [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	result := multiplyMatrices(matrixA, matrixB)

	fmt.Println("Resultant Matrix:")
	for _, row := range result {
		fmt.Println(row)
	}
}
