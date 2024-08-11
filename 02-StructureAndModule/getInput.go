package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	names()
}

func names() {
	fmt.Println("Enter your name:")

	// Using bufio to handle full line input, including spaces
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Read input
	name := scanner.Text()

	// Check if the user entered an empty string
	if len(name) == 0 {
		fmt.Println("You didn't enter a name.")
		return
	}

	// Check whether the name contains a vowel
	for _, v := range strings.ToLower(name) {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			fmt.Println("Your name contains a vowel.")
			return
		}
	}
	fmt.Println("Your name does not contain a vowel.")
}

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	names()
// }

// func names() {
// 	fmt.Println("Enter your name:")

// 	var name string
// 	fmt.Scanln(&name)
// 	// Check whether name has a vowel
// 	for _, v := range strings.ToLower(name) {
// 		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
// 			fmt.Println("Your name contains a vowel.")
// 			return
// 		}
// 	}
// 	fmt.Println("Your name does not contain a vowel.")
// }
