package difr

import (
	"math"
)

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

// Computing Discrete Frechet DiscreteFrechetDistance
// http://www.kr.tuwien.ac.at/staff/eiter/et-archive/cdtr9464.pdf
func (df *discreteFrechet) DiscreteFrechetDistance() float64 {

	i := len(df.firstCurve)
	j := len(df.secondCurve)

	if i == 0 || j == 0 {
		return 0
	}

	df.memo = NewMatrix(i, j, -1)

	d := df.discreteFrechetDistance(i-1, j-1)

	return d
}

// On Map-Matching Vehicle Tracking Data
// https://www.researchgate.net/publication/221310236_On_Map-Matching_Vehicle_Tracking_Data
func (df *discreteFrechet) AverageDiscreteFrechetDistance() float64 {

	i := len(df.firstCurve)
	j := len(df.secondCurve)

	if i == 0 || j == 0 {
		return 0
	}

	df.memo = NewMatrix(i, j, -1)

	d := df.averageDiscreteFrechetDistance(i-1, j-1, i, j)

	return d

}

// Curve Matching, Time Warping, and Light Fields: New Algorithms for Computing Similarity between Curves
// https://citeseerx.ist.psu.edu/document?repid=rep1&type=pdf&doi=3b73b0e130945340c58c7f47235a93a0c6094907
func (df *discreteFrechet) DynamicTimeWrapping() float64 {

	i := len(df.firstCurve)
	j := len(df.secondCurve)

	if i == 0 || j == 0 {
		return 0
	}

	df.memo = NewMatrix(i, j, -1)

	d := df.dynamicTimeWrapping(i-1, j-1)

	return d
}

func (df *discreteFrechet) discreteFrechetDistance(i int, j int) float64 {

	if c := df.memo[i][j]; c > -1 {
		return c
	}

	d := df.firstCurve[i].Distance(df.secondCurve[j])

	if i == 0 && j == 0 {
		df.memo[i][j] = d
		return d
	}

	if j == 0 {
		res := math.Max(df.discreteFrechetDistance(i-1, j), d)
		df.memo[i][j] = res
		return res
	}

	if i == 0 {
		res := math.Max(df.discreteFrechetDistance(i, j-1), d)
		df.memo[i][j] = res
		return res
	}

	res := math.Max(math.Min(df.discreteFrechetDistance(i-1, j), math.Min(df.discreteFrechetDistance(i-1, j-1), df.discreteFrechetDistance(i, j-1))), d)
	df.memo[i][j] = res
	return res
}

func (df *discreteFrechet) dynamicTimeWrapping(i int, j int) float64 {

	if c := df.memo[i][j]; c > -1 {
		return c
	}

	d := df.firstCurve[i].Distance(df.secondCurve[j])

	if i == 0 && j == 0 {
		df.memo[i][j] = d
		return d
	}

	if j == 0 {
		res := df.dynamicTimeWrapping(i-1, j) + d
		df.memo[i][j] = res
		return res
	}

	if i == 0 {
		res := df.dynamicTimeWrapping(i, j-1) + d
		df.memo[i][j] = res
		return res
	}

	minDistance := df.dynamicTimeWrapping(i-1, j-1)

	if c := df.dynamicTimeWrapping(i, j-1); c < minDistance {
		minDistance = c
	}

	if c := df.dynamicTimeWrapping(i-1, j); c < minDistance {
		minDistance = c
	}

	res := minDistance + d
	df.memo[i][j] = res
	return res
}

func (df *discreteFrechet) averageDiscreteFrechetDistance(i int, j int, alpha_n int, beta_n int) float64 {

	if c := df.memo[i][j]; c > -1 {
		return c
	}

	if i == 0 || j == 0 {
		df.memo[i][j] = 0
		return 0
	}

	d := df.firstCurve[i].Distance(df.secondCurve[j])

	minDistance := df.averageDiscreteFrechetDistance(i-1, j-1, alpha_n, beta_n)
	coefficient := math.Sqrt(1/math.Pow(float64(alpha_n), 2) + 1/math.Pow(float64(beta_n), 2))

	tmpMinDistance := df.dynamicTimeWrapping(i, j-1)
	tmpCoefficient := 1 / float64(alpha_n)
	if tmpMinDistance+d*tmpCoefficient < minDistance+d*coefficient {
		minDistance = tmpMinDistance
		coefficient = tmpCoefficient
	}

	tmpMinDistance = df.dynamicTimeWrapping(i-1, j)
	tmpCoefficient = 1 / float64(beta_n)
	if tmpMinDistance+d*tmpCoefficient < minDistance+d*coefficient {
		minDistance = tmpMinDistance
		coefficient = tmpCoefficient
	}

	res := minDistance + d*coefficient
	df.memo[i][j] = res
	return res
}
