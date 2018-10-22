package main

import (
	"fmt"
	"os"
	"strconv"
)

type Feet float64
type Meters float64

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		for _, arg := range args {
			argFloat, err := strconv.ParseFloat(arg, 64)
			if err == nil {
				fmt.Println(metersToFeet(argFloat))
				fmt.Println(feetToMeters(argFloat))
			} else {
				fmt.Println("Invalid input")
				os.Exit(1)
			}
		}
	} else {
		var arg float64
		fmt.Scan(&arg)
		fmt.Println(metersToFeet(arg))
		fmt.Println(feetToMeters(arg))

	}
}

func (f Feet) String() string   { return fmt.Sprintf("%g Feet", f) }
func (m Meters) String() string { return fmt.Sprintf("%g Meters", m) }

func metersToFeet(m float64) Feet   { return Feet(m * 3.281) }
func feetToMeters(f float64) Meters { return Meters(f / 3.281) }
