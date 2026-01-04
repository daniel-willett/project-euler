package main
import ("fmt";"strconv")

func reverse(nstr string) string{
	var result string
	for counter := len(nstr)-1; counter>=0; counter--{
		result += string(nstr[counter])
	}
	return result
}

func isPalendrome(n int) bool{
	var nstr string = strconv.Itoa(n)
	if reverse(nstr)==string(nstr){
		return true
	}
	return false
}

func max(counter []int) int{
	var largest int = 0
	for _, val := range counter{
		if val>largest{
			largest=val
		}
	}
	return largest
}

func main() {
	counter := []int{}
	var n int = 0
	for i:=999; i>100; i--{
		for j:=999; j>100; j--{
			n=i*j
			if isPalendrome(n)==true{
				counter = append(counter,n)
			}
		}
	}
	fmt.Println(max(counter))
}
