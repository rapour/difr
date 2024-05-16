package difr

func CastToPoints[T Point](humans []T) []Point {
	result := []Point{}
	for _, h := range humans {
		result = append(result, h)
	}
	return result
}
