package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("The program called is: ", os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
