package main
import ("fmt"; "math")

func isPrime(n int) bool{
	upper := int(math.Sqrt(float64(n)))+1
	for factor:=2; factor<upper; factor++{
		if n%factor==0{
			return false
		}
	}
	return true
}

func main() {
	var target int = 600851475143

	for i:=2; i<=target; i++{
		if target%i==0{
			if isPrime(target/i)==true{
				fmt.Println(target/i)
				break
			}
		}
	}
}
