package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	fmt.Scan(&r)

	luas := math.Pi * r * r
	keliling := 2 * math.Pi * r

	fmt.Println(luas)
	fmt.Println(keliling)

}
