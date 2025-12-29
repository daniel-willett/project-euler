package main
import("fmt")

func collatz(n int) int{
	var counter int = 0
	for n>1 {
		if n%2==0{
			n = n/2
		} else {
			n = (3*n) + 1
		}
		counter += 1
	}
	return counter
}

func main() {
	var runningN int = 0
	var runningCollatz int = 0
	var n int = 1
	for n<1000000 {
		var value int = collatz(n)
		if value>runningCollatz{
			runningCollatz = value
			runningN = n
		}
		n += 1
	}
	fmt.Println(runningN)
}
