package main

import (
	"fmt"
)

func main(){

	// Part 1 Primmitive data
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

	// Part 2 - pointers
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

	// Part 3 - address of operator - pointer
	favoriteCity := "Helsinki"
	fmt.Println(favoriteCity)

	ptr := &favoriteCity
	fmt.Println(ptr, *ptr)

	favoriteCity = "Oslo"
	fmt.Println(ptr, *ptr)

	// same memory spot
	// 0xc000042260 Helsinki
	// 0xc000042260 Oslo


}
