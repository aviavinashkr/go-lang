// import main package every time
package main

// import package
import (
	"fmt"
)

var test int

func main() {
	var name string = "Avinash"
	var age int = 15
	var phone float32 = 99.99
	var bool bool = true
	test = 30

	//":=" this will decide the data type on basics of values passed to variable
	a := "tesing"

	//contant variable
	const PI = 3.14

	//array
	// array has definite size
	var arr = [3]int{4, 5, 6}
	arr1 := [4]int{7, 8, 9, 10}

	// "[...]" this will define arry is indefinte size
	arr2 := [...]int{1, 2, 3, 4, 5, 6, 2, 1, 3, 1, 4}

	// go slice are similar to array , but more powerful and flexible. It can be used to store multiple values of the same type in a single variable, length of a slice can grow and shrink as see fit.
	arr_slice := []int{1, 2}

	//len() and cap() function are use to return the lenght and capacity of a slice
	fmt.Println(len(arr_slice))
	fmt.Println(cap(arr_slice))

	var add = 100 + 88
	fmt.Println(add)
	fmt.Println(arr_slice)
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println(PI)
	var x, y, z int = 1, 2, 3
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(test)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(phone)
	fmt.Println(bool)

	fmt.Println("hello world!") // this is a command

	//if else statment
	if age > 18 {
		fmt.Println("You are an adult.")
	} else if age == 15 {
		fmt.Println("You are a minor.")
	} else {
		fmt.Println("You are not an adult.")
	}

	//nested if statement
	if age < 18 {
		fmt.Println("You can vote.")
		if bool {
			fmt.Println("You can drive.")
		}
	}

	//switch statment
	switch age < 18 {
	case age == 15:
		fmt.Println("You are a teenager.")
	case age == 18:
		fmt.Println("You are an adult")
	default:
		fmt.Println("You are not an adult.")
	}

	//for loop
	for i := 0; i <= 9; i++ {
		fmt.Printf("%d ", i)
	}

	newfunction()
	functionWithParam(19)
	counters(1)
	var p1 person
	p1.name = "avinash"
	p1.age = 21
	var p2 person
	p2.name = "kumar"
	p2.age = 24
	fmt.Printf(p1.name)
	fmt.Print(p2.name)

	//maps in go
	var maps = map[string]string{"name": "avinash", "age": "21"}
	fmt.Print(maps)

	//empty map in go
	// var amap map[string]string
	// var bmap = make(map[string]string)

}

// function creation
func newfunction() {
	fmt.Println("This function has been called ")
}

func functionWithParam(age int) {
	fmt.Println("This function with param has been called", age)
}

// resusion in go
func counters(x int) int {
	if x == 12 {
		return 0
	}
	fmt.Println(x)
	return counters(x + 1)
}

// structures in go
type person struct {
	name    string
	age     int
	rollno  int
	address string
}
