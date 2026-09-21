package main

import (
	"fmt"
	"math"
)

func factor(n int) [][]int{
	primes := []int{2,3,5,7,11,13,17,19}
	result := [][]int{}
	var cannotDivide bool = false
	var numOfDivisors int = 0
	for _, p := range primes{
		cannotDivide = false
		numOfDivisors = 0
		for cannotDivide==false{
			if n%p!=0{
				cannotDivide = true
			} else {
				numOfDivisors += 1
				n = n/p
			}
		}
		result = append(result, []int{p,numOfDivisors})
	}
	return result
}

func main(){
	counter := map[int]int{2:0,
				3:0,
				5:0,
				7:0,
				11:0,
				13:0,
				17:0,
				19:0}
	for i:=2; i<=20; i++{
		factors := factor(i)
		for _, valuePair := range factors{
			prime := valuePair[0]
			if counter[prime]<valuePair[1]{
				counter[prime] = valuePair[1]
			}
		}
	}
	var total int = 1
	for prime, power := range counter{
		total *= int(math.Pow(float64(prime), float64(power)))
	}
	fmt.Println(total)
}
