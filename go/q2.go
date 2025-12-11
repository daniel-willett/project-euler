 package main
import ("fmt")

func main() {
	arr := []int{1,2}
	var counter int = 1
	var total int = 0
	for arr[counter]<=4000000 {
		if arr[counter]%2==0{
			total += arr[counter]
		}
		arr = append(arr, arr[counter]+arr[counter-1])
		counter+=1
	}
	fmt.Println(total)
}
