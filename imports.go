package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now your have %g problems.", math.Nextafter(2, 3))
}