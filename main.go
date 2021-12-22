package main

import (
	"fmt"
)

const language = "ENGLISH"

const (
	city1 = "Los Angeles"
	num10 = iota
	num11 = iota
	num12 = iota + 6
	num13 = 3 << iota
	num14 = 10 << iota
	num15
)

const (
	num16 = iota
	num17
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

	// part 4 - constants
	const pie = 3.123234
	fmt.Println(pie)

	const num1 = 9
	fmt.Println(num1 + 4)

	// bunch of code
	fmt.Println(num1 + 3.2)


	// differnt data type
	const num2 int = 34
	fmt.Println ( num2 + 10)

	const num3 int = 100
	fmt.Println ( float32(num3) + 3.3)



// accessing outside variable
fmt.Println(language)

// accessing iota (los angeles, 1, 2)
fmt.Println(city1, num10, num11, num12, num13)

// accessing iota (320, 640)
fmt.Println(num14, num15)

// accessing iota (320, 640, 0)
// iota will reset with a new constant block
fmt.Println(num14, num15, num16)

// accessing iota (0, 1)
fmt.Println(num16, num17)
}
