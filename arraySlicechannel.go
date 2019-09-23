package arraySlicechannel

import (
	"fmt"
)

func arraySlicechannel() {
	//var a [5]int = [5]int{1,2,3,4,5}
	a := [5]int{1, 2, 3, 4, 5}
	myslice := []int{1, 2, 3, 4, 5}
	myslice = append(myslice, 6, 7, 8)
	fmt.Println(myslice)
	fmt.Println(a)
	s := myslice[2:4]
	fmt.Println(s[0:3])
	buffch := make(chan int, 2)
	buffch <- 3
	buffch <- 2
	fmt.Println(<-buffch)
}
