package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("The program called is: ", os.Args[0])

	for index, value := range os.Args[1:] {
		fmt.Println("Argument index:", index+1)
		fmt.Println("Argument value:", value)
	}
	//fmt.Println(strings.Join(os.Args[1:], " "))
}
