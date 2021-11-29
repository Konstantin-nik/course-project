package main

import (
	"fmt"
	"math"
)

func f(x float64, c2 chan float64) {
	c2 <- float64(math.Cos(math.Pow(x, 2)) / (1.0 + math.Pow(math.Log(x+1), 2)))
}

func someFunc(p int, c1 chan int, c2 chan float64) {
	i := <-c1
	go f(2*(math.Cos(float64(2*i+1)*math.Pi/float64(2*p+2)))+4, c2)
}

func fHaussa(p int, l chan float64) {
	var ans float64 = 0
	ch := make(chan int)
	ca := make(chan float64)
	for i := 0; i < p; i++ {
		go someFunc(p, ch, ca)
		ch <- i
		ans += <-ca
	}
	l <- ans * math.Pi / float64(p)
}

func main() {
	l := make(chan float64)
	go fHaussa(203207, l)
	//fHaussa(3)

	//fmt.Scan(&x)
	fmt.Print(<-l)
}
