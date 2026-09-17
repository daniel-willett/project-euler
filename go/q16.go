package main
import (
	"fmt"
	"github.com/daniel-willett/common-go-methods"
	"strconv"
)

func main(){
	var num string = "1"
	for i:=1; i<=1000; i++{ //2^1000
		num, _ = common.Multiplication(num,"2")
	}

	var counter int = 0
	var temp int = 0
	for _, val := range num{
		temp, _ = strconv.Atoi(string(val))
		counter += temp
	}
	fmt.Println(counter)
}
