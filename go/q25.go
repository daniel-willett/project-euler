package main

import(
	"fmt"
	"github.com/daniel-willett/common-go-methods"
)

func main(){
	var F1 string = "1"
	var F2 string = "1"
	var temp string = ""
	var index int = 2
	for len(F2)<1000{
		temp, _ = common.Addition(F1,F2)
		F1 = F2
		F2 = temp
		index += 1
	}
	fmt.Println(index)
}
