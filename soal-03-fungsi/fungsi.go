package main
import "fmt"

func main () {
  var x,y,rumus float64
	
	fmt.Scan(&x, &y)
	
	rumus = 1 / (3 * (x * x) + 10) + (10 * y) + 7
	
	fmt.Println(rumus)
}
