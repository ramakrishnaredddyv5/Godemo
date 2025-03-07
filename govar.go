// variable assigning
// there are mainly two methods to declare a variable in go
// 1.by using 'var' keyword and assign a value:-
package main

import "fmt"

func main() {
	var roll_num int
	roll_num = 25
	fmt.Println("my roll number is:", roll_num)
	var name string = "i'm souji"
	fmt.Println("my name is:", name)

	//2.by using short hand method:
	x := 25
	z := "Souji"
	fmt.Println("my roll number is:", x)
	fmt.Println("my name is:", z)

	//3.by declaring multipule values:-
	var c, d int = 5, 10
	fmt.Println("marks are:", c, ",", d)
	a, b := "Souji", "ram"
	fmt.Println("the names are:", a, " ,", b)

	//4.by using composite data types:-
	var arr [5]int = [5]int{80, 70, 89, 90, 95}
	fmt.Println("the marks of a particular student is:=", arr)
}
