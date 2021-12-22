package main

import (
	"fmt"
)

func main(){
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 23.23
	fmt.Println(f)

	//implicit declaration variable
	firstName := "Arthur"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}