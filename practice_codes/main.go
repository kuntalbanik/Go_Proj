//	--
//
// Simple program that prints 'Hello World'
// --

package main

import (
	"fmt"
	"mylearning/myutil"
	userinput "mylearning/userInput"
)

func main() {
	fmt.Println("Hello World")
	myutil.PrintMessage("This message from another file")

	var name string = "John"
	fmt.Println(name)

	var version = "1.0"
	fmt.Println(version)

	// constant variable
	const __software_version__ = "2.0"
	fmt.Println("Software Version :", __software_version__)

	fmt.Println(myutil.PersonName)

	// userinput.UserInput()

	userinput.UserInputBuff()
}
