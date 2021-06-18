package main

import (
	"fmt"
)

func main() {
	var floating_input float64;
	fmt.Println("Enter a decimal number:");
	fmt.Scan(&floating_input);
	var truncat int;
	truncat = int(floating_input);
	fmt.Println(truncat);

}
