package main
import("fmt";"math")

func isPrime(n int) bool{
	upper := int(math.Sqrt(float64(n)))+1
	for factor:=2; factor<upper; factor++{
		if n%factor==0{
			return false
		}
	}
	return true
}

func main(){
	var total int = 0
	for n:=2; n<2000000; n++{
		if isPrime(n)==true{
			total+=n
		}
	}
	fmt.Println(total)
}
