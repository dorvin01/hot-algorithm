package matrix

func setZeroes(matrix [][]int) {
	column := len(matrix)
	row := len(matrix[0])
	isFirstColZero := false
	for y := 0; y < row; y++ {
		if matrix[0][y] == 0 {
			isFirstColZero = true
			break
		}
	}
	isFirstRowZero := false
	for x := 0; x < column; x++ {
		if matrix[x][0] == 0 {
			isFirstRowZero = true
			break
		}
	}

	for x := 1; x < column; x++ {
		for y := 1; y < row; y++ {
			if matrix[x][y] == 0 {
				matrix[x][0] = 0
				matrix[0][y] = 0
			}
		}
	}

	for x := 1; x < column; x++ {
		for y := 1; y < row; y++ {
			if matrix[x][0] == 0 || matrix[0][y] == 0 {
				matrix[x][y] = 0
			}
		}
	}

	if isFirstColZero {
		for y := 0; y < row; y++ {
			matrix[0][y] = 0
		}
	}
	if isFirstRowZero {
		for x := 0; x < column; x++ {
			matrix[x][0] = 0
		}
	}
}
