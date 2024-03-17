package main

import "fmt"
import "math"

func factorial(n float64) float64 {
	// n! = sqrt(2*pi*n)* n^n * exp^-n
	return math.Sqrt(2 * math.Pi * n) * math.Pow(n/math.E,n)
	}
	
func main(){
	var n int 
	fmt.Println("Enter number")
	fmt.Scan(&n)
	
	result := factorial(float64(n))
		fmt.Printf("%d! = %.2f\n",n, result)
}
