package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var operation string

	fmt.Println("select any operation")

	fmt.Scanln(&operation)

	var num1, num2 string

	fmt.Println("enter the first number")
	_, err := fmt.Scanln(&num1)

	if err != nil {

		log.Fatal(err)
	}

	switch operation {
	case "+":
		a := StringtoInt(num1)
		b := StringtoInt(num2)

		ans := sum(a, b)

		fmt.Println("Sum is :", ans)

	case "-":
		a := StringtoInt(num1)
		b := StringtoInt(num2)

		ans := sub(a, b)

		fmt.Println("final ans is :", ans)

	case "/":
		a := Stringtofloat(num1)
		b := Stringtofloat(num2)

		ans := divide(a, b)

		fmt.Println("final ans is :", ans)

	case "*":
		a := StringtoInt(num1)
		b := StringtoInt(num2)

		ans := mul(a, b)

		fmt.Println("ans is :", ans)

	}

}

func StringtoInt(str string) int {

	a, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}

	return a

}

func Stringtofloat(str string) float64 {

	a, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}

	return a

}

func sum(a, b int) int { //Addition Function

	return a + b
}

func sub(a, b int) int { //Substraction function
	return a - b
}

func divide(a, b float64) float64 { // division function

	return a / b
}

func mul(a, b int) int { // multiplying function
	return a * b

}
