package main // Every program will have a package

import "fmt"

func main() { // Every go program has a main function. There can be only one main function in a go program. It is the entry point in the program. func keyword helps us to declare a function in go
	fmt.Println("Hello World !") //fmt stands for format. It is a standard Library
	var whattosay string         //Declaring a string variable
	whattosay = "Goodbye world. Ending the program"
	fmt.Println(whattosaygo)
	var i int //declaring a integer variable
	i = 34
	fmt.Println(i)
	check := say()                               //you are using a variable to call the function say. When you use a variable to call a function you don't declare it.
	fmt.Println("The function returned ", check) //When you want to print something with the variable
	check2, check3 := say2()                     //Since the function say2 is returning 2 things, we need to declare 2 variables here. Hence we have defined check2 and check3
	fmt.Println(check2)
	fmt.Println(check3)

}

func say() string { //This function has something after the paranthesis. The return type is string
	return "Nandita"
}

func say2() (string, t) { // When you want to return multiple things from a func, you use (). enter the variable type of what you want to return in ()
	return "Tankha", 2
}
