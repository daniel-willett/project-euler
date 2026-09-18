package main
import("fmt";"github.com/daniel-willett/common-go-methods")

func main(){
	var total int = 0
	for n:=2; n<2000000; n++{
		if common.IsPrime(n)==true{
			total+=n
		}
	}
	fmt.Println(total)
}
