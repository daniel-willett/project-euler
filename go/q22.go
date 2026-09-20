package main

import(
	"fmt"
	"os"
	"github.com/daniel-willett/common-go-methods"
)

func ALessThanB(A string, B string) bool{
	var index int = 0
	var Alen int = len(A)
	var Blen int = len(B)
	//Some smart logic here, if index goes out of bounds we end the for before executing the part `A[index]==B[index]` which would cause an error
	for index<Alen && index<Blen && A[index]==B[index]{
		index += 1
	}
	if index==Alen || index==Blen{
		if Alen<Blen{
			return true
		} else {
			return false
		}
	}
	a := A[index]
	b := B[index]
	if a<b{
		return true
	} else{
		return false
	}
}

func sort(arr []string) []string{
	//Bubble Sort is sufficient here
	var length int = len(arr)
	var changesMade bool = true
	for changesMade==true{
		changesMade = false
		for index:=0; index<length-1; index++{
			if ALessThanB(arr[index+1],arr[index])==true { //arr[index+1]<arr[index]
				changesMade = true
				temp := arr[index]
				arr[index] = arr[index+1]
				arr[index+1] = temp
			}
		}
	}
	return arr
}

func productOfChars(str string) int{
	var counter int = 0
	for _, val := range str{
		counter += int(val)-64
	}
	return counter
}

func main(){
	data, err := os.ReadFile("q22_names.txt")
	if err != nil{
		panic(err)
	}
	
	input := string(data)
	input = common.Replace(input, "\"", "")
	inputArr := common.Split(input, ",")
	
	sortedArr := sort(inputArr)

	var result int = 0
	for i:=0; i<len(sortedArr); i++{
		result += productOfChars(sortedArr[i])*(i+1)
	}
	fmt.Println(result)
}
