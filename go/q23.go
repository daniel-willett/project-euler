package main
import ("fmt";"math")

func divisors(n int) []int{
	lst := []int{1}
	upper := int(math.Sqrt(float64(n)))
	for i:=2; i<=upper; i++{
		if n%i==0{
			lst = append(lst, i)
			lst = append(lst,n/i)
		}
	}
	if int(math.Sqrt(float64(n)))==n{
		lst = append(lst,int(math.Sqrt(float64(n))))
	}
	return lst
}

func isAbundant(n int) bool{
	divs := divisors(n)
	total := 0
	for i := 0; i<len(divs); i++{
		total += divs[i]
	}
	if total>n{
		return true
	}
	return false
}

func containedIn(i int, abundantSum []int) bool{
	for k:=0; k<len(abundantSum); k++{
		if abundantSum[k]==i{
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Program Start")
	abundantNumbers := []int{}
	for i:=12; i<=28123; i++{
		if isAbundant(i)==true{
			abundantNumbers = append(abundantNumbers,i)
		}
	}
	fmt.Println("abundantNumbers")
	abundantSum := []int{}
	for i:=0; i<len(abundantNumbers); i++{
		for j:=0; j<len(abundantNumbers); j++{
			abundantSum = append(abundantSum, abundantNumbers[i]+abundantNumbers[j])
		}
	}
	fmt.Println("abundantNumbers")
	fmt.Println(len(abundantNumbers))
	total := 0
	for i:=12; i<=28123; i++{
		if containedIn(i,abundantSum)==false{
			total += i
		}
	}
	fmt.Println(total)
}
