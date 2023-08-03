package main

import (
	"fmt"
	"math"
)

const pi = math.Pi

func findPi(val int) (float64, error) {
	ratio := math.Pow(10, float64(val))
	if val > 15 {
		return 0, fmt.Errorf("value must be less than 15")
	}
	return math.Round(pi*ratio) / ratio, nil
}

func main() {
	fmt.Println(findPi(20))
}
