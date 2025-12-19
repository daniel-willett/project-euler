package main
import("fmt"; "math")

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
	var counter int = 2
	var num int = 4
	for counter<10001{
		if isPrime(num)==true{
			counter += 1
		}
		num += 1
	}
	fmt.Println(num-1)
}

