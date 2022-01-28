package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix Define the Matrix type here.
type Matrix struct {
	data [][]int
}

// New create new instance of matrix
func New(s string) (*Matrix, error) {
	var m = Matrix{}
	var splittedStrings = strings.Split(s, "\n")
	var lenOfInnerStr = len(strings.Split(strings.Trim(splittedStrings[0], " "), " "))

	for _, outStr := range splittedStrings {
		var currentRow []int

		splittedInnerStrings := strings.Split(strings.Trim(outStr, " "), " ")

		if len(splittedInnerStrings) != lenOfInnerStr {
			return &Matrix{}, errors.New("uneven amount of columns")
		}

		for _, inStr := range splittedInnerStrings {
			integer, err := strconv.Atoi(inStr)
			if err != nil {
				return &Matrix{}, err
			}

			currentRow = append(currentRow, integer)
		}
		m.data = append(m.data, currentRow)
	}

	return &m, nil
}

// Rows must return the results without affecting the matrix.
func (m *Matrix) Rows() [][]int {
	rowCopy := make([][]int, len(m.data))
	for i := range m.data {
		rowCopy[i] = make([]int, len(m.data[i]))
		copy(rowCopy[i], m.data[i])
	}

	return rowCopy
}

// Cols return columns in a matrix
func (m *Matrix) Cols() [][]int {
	colLength := len(m.data[0])
	var colCopy [][]int

	for idx := 0; idx < colLength; idx++ {
		var col []int

		for _, d := range m.data {
			col = append(col, d[idx])
		}

		colCopy = append(colCopy, col)
	}

	return colCopy
}

// Set set value at an index intersection in the matrix
func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m.data) {
		return false
	}

	if col < 0 || col >= len(m.data[0]) {
		return false
	}

	m.data[row][col] = val
	return true
}
