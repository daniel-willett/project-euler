package main
import("fmt"; "math")

func numOfDivs(x int) int{
	var counter int = 0
	var upper int = int(math.Sqrt(float64(x)))+1
	for factor:=1; factor<upper; factor++{
		if x%factor==0{
			counter+=1
		}
	}
	counter*=2
	if float64(x)==math.Sqrt(float64(x)){
		counter-=1
	}
	return counter
}

func main(){
	var n int = 1
	var triangle int = 1
	for numOfDivs(triangle)<=500{
		n+=1
		triangle=n*(n+1)/2
	}
	fmt.Println(triangle)
}
