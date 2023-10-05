package main

import "fmt"

var x int = 1

//y := 1 //error declaration

func main() {
	var myName string
	myName = "Bariq Firjatullah"
	var nickName = "Bariq"
	myAge := 22

	var a string
	var b int
	var c bool

	var d, e, f, g = 1, 2, 3, "Hello World"

	var (
		firstName string = "Bariq"
		lastName         = "Firjatullah"
		height    int    = 170
	)

	fmt.Println(myName)
	fmt.Println(nickName)
	fmt.Println(myAge)

	fmt.Println(x)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(height)
}
