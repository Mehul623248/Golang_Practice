package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Enter a string:")
	fmt.Scan(&input)
	if strings.HasPrefix(input, "I") || strings.HasPrefix(input, "i") {
		if strings.HasSuffix(input, "n") || strings.HasSuffix(input, "N") {
            if strings.Contains(input, "a") || strings.Contains(input, "A"){
                fmt.Println("Found!");
			}else{
				fmt.Println("Not Found!");
			}
			
		} else {
			fmt.Println("Not Found!");
		}
	} else {
		fmt.Println("Not Found!");
	}
	//

}
