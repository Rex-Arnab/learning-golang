package main

import (
	"bufio"
	"fmt"
	"os"
)

func helloWorld(){
	fmt.Println()
	fmt.Println("Hello world")
}

func commandLineArgs(){
	fmt.Println()
	fmt.Println("These are Command Line Args : ")
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("The first command line arg is : ", os.Args[1])
	}
}

func environmentVariables(){
	fmt.Println()
	fmt.Println("These are Environment Variables : ")
	fmt.Println(os.Environ())
	fmt.Println()
}

func declaringVariables(){
	fmt.Println()
	fmt.Println("Declaring Variables : ")
	// Method 1
	var (
		name string
		age int
		location string
	)
	name = "Arnab Biswas"
	age = 25
	location = "Kolkata"

	// Method 2
	// var name string = "Arnab Biswas"
	// var age int = 25
	// var location string = "Kolkata"

	// Method 3
	// name := "Arnab Biswas"
	// age := 25
	// location := "Kolkata"
	fmt.Println(name, age, location)
}

func ifElse(){
	fmt.Println()
	fmt.Println("If Else : ")
	// Method 1
	if 5 > 1 {
		fmt.Println("5 is greater than 1")
	}

	// Method 2
	if 5 > 1 {
		fmt.Println("5 is greater than 1")
	} else {
		fmt.Println("5 is not greater than 1")
	}

	// Method 3
	if 5 > 1 {
		fmt.Println("5 is greater than 1")
	} else if 5 == 1 {
		fmt.Println("5 is equal to 1")
	} else {
		fmt.Println("5 is not greater than 1")
	}
}

func switchStatement(){
	fmt.Println()
	fmt.Println("Switch Statement : ")

	num := 5
	switch num {
	case 1:
		fmt.Println("The number is 1")
	case 2:
		fmt.Println("The number is 2")
	case 3:
		fmt.Println("The number is 3")
	case 4:
		fmt.Println("The number is 4")
	case 5:
		fmt.Println("The number is 5")
	default:
		fmt.Println("The number is more then 5")
	}
}

func forLoop(){
	fmt.Println()
	fmt.Println("For Loop")
	for i := 1; i <= 10; i++ {
		// String Formatting
		fmt.Printf("5 X %d = %d\n", i, 5*i)
	}
}

func whileLoop(){
	fmt.Println()
	fmt.Println("While Loop")
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func infiniteLoop(){
	fmt.Println()
	fmt.Println("Infinite Loop")
	k := 1
	for {
		fmt.Println(k)
		k++
		if k == 10 {
			break
		}
	}
}

func arrayMethods(){
	// 	// Arrays
	fmt.Println("Arrays")
	// Method 1
	// var arr [5]int
	// arr[0] = 10
	// arr[1] = 20
	// arr[2] = 30
	// arr[3] = 40
	// arr[4] = 50

	// Method 2
	var arr = [5]int{10, 20, 30, 40, 50}

	// Method 3
	// var arr = [...]int{10, 20, 30, 40, 50}

	fmt.Println(arr)
}

func slices(){
	// Slices
	fmt.Println("Slices")
	// Method 1
	// var slice []int
	// slice = append(slice, 10)
	// slice = append(slice, 20)
	// slice = append(slice, 30)
	// slice = append(slice, 40)
	// slice = append(slice, 50)

	// Method 2
	// var slice = []int{10, 20, 30, 40, 50}

	// Method 3
	var slice = [...]int{10, 20, 30, 40, 50}

	fmt.Println(slice)
}

func maps(){
	fmt.Println("Maps")
	var person = map[string]int{
		"Rahul":  20,
		"Rohit":  30,
		"Rajesh": 40,
	}
	fmt.Println(person)
}

func inputFromKeyboard(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Name : ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Hello, "+ name)
}

func structExample(){
	fmt.Println("Structs")
	type Person struct {
		name string
		age int
		location string
	}
	var person Person
	person.name = "Arnab Biswas"
	person.age = 25
	person.location = "Kolkata"
	fmt.Println(person)

	// Diff between struct and variable
	// Struct is a collection of variables
	// Variable is a single variable

}

func interfaceExample(){
	fmt.Println("Interfaces")


}

func main(){
	helloWorld()
	commandLineArgs()
	environmentVariables()
	declaringVariables()
	ifElse()
	switchStatement()
	forLoop()
	whileLoop()
	infiniteLoop()
	// Diff between Array and Slices
	// Array is fixed length
	// Slices are dynamic
	arrayMethods()
	slices()
	maps()
	inputFromKeyboard()
	structExample()
}