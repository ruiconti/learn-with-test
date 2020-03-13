package arrays

const defaultSize = 2

func CreateMatrix(size int) (matrix [][]float64) {
	if size == 0 {
		size = defaultSize
	}
	matrix = make([][]float64, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]float64, size)
	}
	return
}

func CreateIdentityMatrix(size int) (matrix [][]float64) {
	matrix = CreateMatrix(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				matrix[i][j] = 1
			}
		}
	}
	return
}
