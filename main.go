// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

//var x uint8 = 3

//y := 5
func main() {
	fmt.Println("Hello from Kafee")
	fmt.Println(CompareNumbers(2, 4))
	//x := 2
	//var x uint8 = 2
	//fmt.Println(x)
	switch x := 2; x {
	case 3:
		fmt.Println("I am 3")
		fallthrough
	case 2: 
		fmt.Println("I am 2")
	case 4:
		fmt.Println("I am 4")
	}
	for i :=1;i < 40;i++{
		fmt.Println(i)
	}
}

func CompareNumbers(i1, i2 int) (bool, int) {
	/*
		if (i1 > i2){
			fmt.Println("First number is greater than the second number")
			return false, i1-i2
		} else if(i2 > i1){
			fmt.Println("Second number is greater than the First number")
			return false, i2-i1
		} else{
			return true, 0
		}
	*/
	switch {
	case i1 > i2:
		fmt.Println("First number is greater than the second number")
		return false, i1 - i2
	case i2 > i1:
		fmt.Println("Second number is greater than the First number")
		return false, i2 - i1
	}
	return true, 0
}