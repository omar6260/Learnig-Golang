
package main

import "fmt"

func main() {
	// strings 
	// var nameOne string = "Oumar"
	// var nameTwo = "Saliou"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "peach"
	// nameThree = "bowser"

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "yoshi"
	// fmt.Println(nameFour)

	// ints

	// var ageOne int = 10
	// var ageTwo = 20
	// ageThree := 40

	// fmt.Println(ageOne, ageTwo, ageThree)

	//  bits & memory
	// var nameOne int8 = 25
	// var nameTwo int8 = -200
	// var nameThree uint16 = 256

	// var nameOne string = "Fall"
	// var prenom = "Oumar"
	// var ageOne int = 24
	// date := "ans"

	// fmt.Println(prenom, nameOne, ageOne, date)

	// Fonction Print

	name := "Oumar"
	age := 24
	var i interface{} = 23

	// Print
	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	fmt.Print("we are a new line \n")

	// Println
	fmt.Println("hello oumar")
	fmt.Println("Goodbye Omar")
	fmt.Println("my name is", name, "and my age is", age)

	// Printf (formated strings) %v _ = format specifier
	fmt.Printf("my name is %v and my age is %v \n", name, age)
	fmt.Printf("age is of type %T \n")
	fmt.Printf("you scrored %f points! \n", 222.222)
	fmt.Printf("you scrored %0.1f points! \n", 222.222)
	fmt.Printf("%v\n", i)
	fmt.Printf("%[2]d %[1]d\n", 11, 22)

	// Sprintf (save formated str)

	var str = fmt.Sprintf("my name is %v and my age is %v \n", name, age)
	fmt.Println("the save string", str)
	
}