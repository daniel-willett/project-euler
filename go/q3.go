package main
import("fmt"; "github.com/daniel-willett/common-go-methods")

func main() {
	var target int = 600851475143

	for i:=2; i<=target; i++{
		if target%i==0{
			if common.IsPrime(target/i)==true{
				fmt.Println(target/i)
				break
			}
		}
	}
}
