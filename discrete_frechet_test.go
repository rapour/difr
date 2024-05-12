package difr

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

type point struct {
	x float64
	y float64
}

func (p point) Distance(op Point) float64 {

	if o, ok := op.(point); ok {
		return math.Sqrt(math.Pow((p.x-o.x), 2) + math.Pow((p.y-o.y), 2))
	}

	return 0
}

var _ Point = point{}

func TestPointDistance(t *testing.T) {

	cases := []struct {
		Name     string
		PointA   Point
		PointB   Point
		Distance float64
	}{
		{
			Name:     "zero distance",
			PointA:   point{x: 1, y: 1},
			PointB:   point{x: 1, y: 1},
			Distance: 0,
		},
		{
			Name:     "1 distance",
			PointA:   point{x: 1, y: 1},
			PointB:   point{x: 0, y: 1},
			Distance: 1,
		},
		{
			Name:     "1 distance second",
			PointA:   point{x: 1, y: 1},
			PointB:   point{x: 1, y: 0},
			Distance: 1,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {

			require.Equal(t, c.Distance, c.PointA.Distance(c.PointB))

		})
	}

}

func TestFrechetDistance(t *testing.T) {

	cases := []struct {
		Name        string
		FirstCurve  []Point
		SecondCurve []Point
		Distance    float64
	}{
		{
			Name:        "zero distance",
			FirstCurve:  []Point{point{x: 1, y: 1}, point{x: 2, y: 2}},
			SecondCurve: []Point{point{x: 1, y: 1}, point{x: 2, y: 2}},
			Distance:    0,
		},
		{
			Name:        "1 distance",
			FirstCurve:  []Point{point{x: 0, y: 0}, point{x: 1, y: 0}},
			SecondCurve: []Point{point{x: 1, y: 0}, point{x: 1, y: 1}},
			Distance:    1,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {

			df := NewDiscreteFrechet(c.FirstCurve, c.SecondCurve)

			require.Equal(t, c.Distance, df.Distance())

		})
	}

}
