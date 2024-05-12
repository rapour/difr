package difr

import "math"

type Point interface {
	Distance(p Point) float64
}

type discreteFrechet struct {
	memo        matrix
	firstCurve  []Point
	secondCurve []Point
}

func NewDiscreteFrechet(firstCurve []Point, secondCurve []Point) *discreteFrechet {

	return &discreteFrechet{
		firstCurve:  firstCurve,
		secondCurve: secondCurve,
	}
}

func (df *discreteFrechet) Distance() float64 {

	i := len(df.firstCurve)
	j := len(df.secondCurve)

	df.memo = NewMatrix(i, j, -1)

	return df.distance(i-1, j-1)
}

func (df *discreteFrechet) distance(i int, j int) float64 {

	if c := df.memo[i][j]; c > -1 {
		return c
	}

	d := df.firstCurve[i].Distance(df.secondCurve[j])

	if i == 0 && j == 0 {
		df.memo[i][j] = d
		return d
	}

	if j == 0 {
		res := math.Max(df.distance(i-1, j), d)
		df.memo[i][j] = res
		return res
	}

	if i == 0 {
		res := math.Max(df.distance(i, j-1), d)
		df.memo[i][j] = res
		return res
	}

	res := math.Max(math.Min(df.distance(i-1, j), math.Min(df.distance(i-1, j-1), df.distance(i, j-1))), d)
	df.memo[i][j] = res
	return res
}
