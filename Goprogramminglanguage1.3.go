package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	// echo2

	s, sep := "", ""

	start_plus := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	elapsed_plus := time.Since(start_plus)
	fmt.Println("Time for for range loop using '+': ", elapsed_plus)

	// echo3

	start_join := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	elapsed_join := time.Since(start_join)
	fmt.Println("Time for 'join': ", elapsed_join)

	//final echo
	start_final := time.Now()
	fmt.Println("The program called is: ", os.Args[0])

	for index, value := range os.Args[1:] {
		fmt.Println("Argument index:", index+1)
		fmt.Println("Argument value:", value)
	}
	elapsed_final := time.Since(start_final)
	fmt.Println("Time for final: ", elapsed_final)
}
