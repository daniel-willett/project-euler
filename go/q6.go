package main
import("fmt";"math")

func main() {
	var n int = 100
	var sumSquared int = int(math.Pow(float64(n*(n+1)/2),float64(2)))
	//fmt.Println(sumSquared)
	var squareSum int = 0
	for i:=1; i<=n; i++{
		squareSum += int(math.Pow(float64(i),float64(2)))
	}
	fmt.Println(sumSquared-squareSum)
}
