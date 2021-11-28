package main

import (
	"fmt"
	"math"
)

var (
	x float64
)

func f(x float64) float64 {
	return float64(math.Cos(math.Pow(x, 2)) / (1.0 + math.Pow(math.Log(x+1), 2)))
}

func someFunc(i int, p int) float64 {
	return f(2*(math.Cos(float64(2*i+1)*math.Pi/float64(2*p+2))) + 4)
}

func fHaussa(p int) float64 {
	var ans float64 = 0
	for i := 0; i < p; i++ {
		//
	}
	return ans * math.Pi / float64(p)
}

func main() {
	go fHaussa(1024)
	go fHaussa(3)

	fmt.Scan(&x)
	fmt.Print(f(x))
}
