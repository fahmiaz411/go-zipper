package zipper

import (
	"math"
)


func Zip(a, b []interface{}) [][]interface{} {
	min := int(math.Min(float64(len(a)), float64(len(b))))
	result := [][]interface{}{}

	for i := 0; i < min; i++ {
		result = append(result, []interface{}{a[i], b[i]})
	}

	return result
}