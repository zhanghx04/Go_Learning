package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result := sum(4, 6)
	fmt.Println(result)

	res, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func sum(x int, y int) int {
	return x + y
}

// error here means sqrt function could print out error
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}
	// nil here is like null but not like null, zero values
	return math.Sqrt(x), nil
}
