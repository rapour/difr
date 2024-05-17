package difr

import (
	"math"
)

type Point interface {
	Distance(p Point) float64
}

type discreteFrechet struct {
	memo        matrix
	weight      matrix
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

	if i == 0 || j == 0 {
		return 0
	}

	df.memo = NewMatrix(i, j, -1)

	d := df.distance(i-1, j-1)

	return d
}

// returns the (distance, weight) ordered pair
func (df *discreteFrechet) DistanceWithWeight() (float64, float64) {

	i := len(df.firstCurve)
	j := len(df.secondCurve)

	if i == 0 || j == 0 {
		return 0, 0
	}

	df.memo = NewMatrix(i, j, -1)
	df.weight = NewMatrix(i, j, 0)

	d := df.distanceWithWeight(i-1, j-1)

	return d, df.weight[i-1][j-1]
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

func (df *discreteFrechet) distanceWithWeight(i int, j int) float64 {

	if c := df.memo[i][j]; c > -1 {
		return c
	}

	d := df.firstCurve[i].Distance(df.secondCurve[j])

	if i == 0 && j == 0 {
		df.memo[i][j] = d
		df.weight[i][j] = d
		return d
	}

	if j == 0 {
		res := math.Max(df.distanceWithWeight(i-1, j), d)
		df.memo[i][j] = res
		df.weight[i][j] = df.weight[i-1][j] + d
		return res
	}

	if i == 0 {
		res := math.Max(df.distanceWithWeight(i, j-1), d)
		df.memo[i][j] = res
		df.weight[i][j] = df.weight[i][j-1] + d
		return res
	}

	minDistance := df.distanceWithWeight(i-1, j-1)
	minWeight := df.weight[i-1][j-1]

	if c := df.distanceWithWeight(i, j-1); c < minDistance {
		minDistance = c
		minWeight = df.weight[i][j-1]
	}

	if c := df.distanceWithWeight(i-1, j); c < minDistance {
		minDistance = c
		minWeight = df.weight[i-1][j]
	}

	res := math.Max(minDistance, d)
	df.memo[i][j] = res
	df.weight[i][j] = minWeight + d
	return res
}
