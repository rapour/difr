package difr

type matrix [][]float64

func NewMatrix(x int, y int, dflt float64) matrix {
	m := make([][]float64, x)
	for rwIndx := range m {

		row := make([]float64, y)

		for clmnIndx := range row {
			row[clmnIndx] = dflt
		}

		m[rwIndx] = row
	}

	return m
}
