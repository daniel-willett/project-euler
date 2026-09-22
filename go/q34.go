package main

import (
	"fmt"
	"github.com/daniel-willett/common-go-methods"
	"strconv"
)

func factSum(n int) int{
	var runningTotal int = 0
	var stringN string = strconv.Itoa(n)
	for _, digit := range stringN{
		temp, _ := common.Factorial(int(digit)-48)
		runningTotal += temp
	}
	return runningTotal
}

func main(){
	var total int = 0
	for i:=3; i<100000; i++{
		if factSum(i)==i{
			total += i
		}
	}
	fmt.Println(total)
}
