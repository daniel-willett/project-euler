package main

import("fmt")

func numberToWord(n int)string{

	var units = map[int]string{0: "",
	1: "one",
	2:"two",
	3:"three",
	4:"four",
	5:"five",
	6:"six",
	7:"seven",
	8:"eight",
	9:"nine",
	10:"ten",
	11:"eleven",
	12:"twelve",
	13:"thirteen",
	14:"fourteen",
	15:"fifteen",
	16:"sixteen",
	17:"seventeen",
	18:"eighteen",
	19:"nineteen"}

	var tens = map[int]string{2:"twenty",
	3:"thirty",
        4:"forty",
        5:"fifty",
        6:"sixty",
        7:"seventy",
        8:"eighty",
        9:"ninety"}

	var hundred string = "hundred"


	var theRest int = n
	var word string = ""

	if int(theRest/100)!=0{
		word += units[int(theRest/100)] + hundred
		theRest = theRest%100
		if theRest!=0{
			word += "and"
		}
	}
	if theRest>19{
		word += tens[int(theRest/10)] + units[theRest%10]
	} else if theRest!=0{
		word += units[theRest]
	}

	return word
}

func main(){
	var str string = ""
	for i:=1; i<1000; i++{
		str += numberToWord(i)
	}
	str += "onethousand"
	fmt.Println(len(str))
}
