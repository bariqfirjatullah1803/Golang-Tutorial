package main

import "fmt"

func main() {
	var firstName, lastName string = "Bariq", "Firjatullah"
	var a, b = 100, 200

	var nickName, myAddress = "Bariq", "Malang"

	var x string = "Halo"
	var y int = 25

	fmt.Print(firstName, " ", lastName, "\n")
	fmt.Print(a, b, "\n")

	fmt.Println(nickName)
	fmt.Println(myAddress)

	fmt.Printf("value x adalah : %v dan tipe data x adalah : %T\n", x, x)
	fmt.Printf("value y adalah : %v dan tipe data x adalah : %T\n", y, y)
}
