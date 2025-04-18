package userinput

import (
	"bufio"
	"fmt"
	"os"
)

func UserInput() {
	var name string
	fmt.Println("Enter your name :")
	fmt.Scan(&name)
	fmt.Println("Hi! ", name)
}

func UserInputBuff() {
	// var name string
	fmt.Println("Enter your name :")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hi! ", name)

}
