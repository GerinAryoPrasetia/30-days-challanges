package main

import (
	"fmt"
	"math"
)

const pi = math.Pi

func findPi(val int) float64 {
	ratio := math.Pow(10, float64(val))
	return math.Round(pi*ratio) / ratio
}

func main() {
	fmt.Println(findPi(10))
}
