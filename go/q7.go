package main
import("fmt"; "github.com/daniel-willett/common-go-methods")

func main(){
	var counter int = 2
	var num int = 4
	for counter<10001{
		if common.IsPrime(num)==true{
			counter += 1
		}
		num += 1
	}
	fmt.Println(num-1)
}

