package main
import("fmt"; "github.com/daniel-willett/common-go-methods")

func main(){
	var n int = 1
	var triangle int = 1
	for common.NumOfDivs(triangle)<=500{
		n+=1
		triangle=n*(n+1)/2
	}
	fmt.Println(triangle)
}
