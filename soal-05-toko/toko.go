package main
import "fmt"

func main () {
	var b1, b2, b3 int
	
	fmt.Scan(&b1, &b2, &b3)
			
	fmt.Println(float64(b1) + (float64(b1) * 5/100), float64(b2) + (float64(b2) * 5/100), float64(b3) + (float64(b3) * 5/100))
}
