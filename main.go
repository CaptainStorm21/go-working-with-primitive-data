package main

import (
	"fmt"
)

func main(){

	// Primmitive data
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

	// Pointers

	/*
		setting up pointers by preceding with a *
		deferennced pointers by preceding the var with an *
	*/
	var santaName *string = new(string)

	// dereferencing
	*santaName = "Santa Claus"

	fmt.Println(santaName)
	// output 0xc000088230 = reference to memory box

	fmt.Println(*santaName)
	// output is a pointer
}