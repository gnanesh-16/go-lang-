/* 	1. func div (num int, deno int) int in here int at the closed bracket defines the end returing with (int)   */


package main

import (
	"errors"
	"fmt"
)

func main() {
	main2()
	num := 11
	deno := 2

	// Capture the result, remainder, and error from the div function
	res, remind, err := div(num, deno)
	
	// Check if there was an error
	if err != nil {
		fmt.Println(err.Error())
	} else if remind == 0 {
		fmt.Println(res)
	} else {
		fmt.Printf("Result of parametrized div is [%v] with remainder [%v]", res, remind)
	}
}

func main2() {
	fmt.Println("$ Utilizing main2-main1 via function call ")
}

// Parametrized div function
func div(num int, deno int) (int, int, error) {
	if deno == 0 {
		err := errors.New("Cannot divide by Zero")
		return 0, 0, err
	}
	res := num / deno
	remind := num % deno
	return res, remind, nil
}
