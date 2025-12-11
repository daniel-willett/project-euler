package main
import ("fmt")

func main() {
	var counter int = 0
	for i:=0; i<1000; i++{
		if (i%3==0 || i%5==0){
			counter += i
		}
	}
	fmt.Println(counter)
}
