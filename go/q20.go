package main

import("fmt";"github.com/daniel-willett/common-go-methods";"strconv")

func main(){
	// common.Factorial doesn't work for this. We'll have to employ the big numbers
	var temp string = "1"
	for i:=1; i<=100; i++{
		temp, _ = common.Multiplication(temp,strconv.Itoa(i))
	}
	var counter string = "0"
	for _, val := range temp{
		counter, _ = common.Addition(counter, string(val))
	}
	fmt.Println(counter)
}
