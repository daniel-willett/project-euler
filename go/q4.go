package main
import("fmt";"github.com/daniel-willett/common-go-methods")

func main() {
	counter := []int{}
	var n int = 0
	for i:=999; i>100; i--{
		for j:=999; j>100; j--{
			n=i*j
			if common.IsPalendrome(n)==true{
				counter = append(counter,n)
			}
		}
	}
	var temp int = 0
	temp, _ = common.Max(counter)
	fmt.Println(temp)
}
