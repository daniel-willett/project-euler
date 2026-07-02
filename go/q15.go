package main

import "fmt"

//Same idea as q15.py
//We do have to be careful about int sizes. So we use float64 as this can go upto 1.7e+308
//40!/(20!*20!) = 21*...*40/20! is a simplification we can make

func factorial(n float64) float64{
	var result float64 = 1
	for i:=n; i>1; i--{
		result *= i
	}
	return result
}

func main(){
	var total float64 = 1
	for i:=float64(21); i<=float64(40); i++{
		total *= i
	}
	top := total
	bottom:= factorial(20)
	fmt.Println(int(top/bottom))
}

//Also coincidently, if we did factorial(40)/(factorial(20)*factorial(20)) then we'd start to get rounding errors bleed through. These would be noticeable of the order 0.00003 but still there noless
