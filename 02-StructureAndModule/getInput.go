// Input Method: The first method (bufio.Scanner) allows for full line input, including spaces, while the second method (fmt.Scanln) captures only the first word.
// Empty Input Handling: The first method checks for empty input before proceeding, whereas the second method does not perform such a check explicitly.
// Use Cases: If you want to allow users to enter full names, method 1 is preferable. If you only need a single word and don’t mind splitting names, method 2 could suffice but has limitations.

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
	name := scanner.Text() // retrieves the text that was read by the scanner.

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
method  2:
package main

import (
	"fmt"
	"strings"
)

func main() {
	names()
}

func names() {
	fmt.Println("Enter your name:")

	var name string
	fmt.Scanln(&name)
	// Check whether name has a vowel
	for _, v := range strings.ToLower(name) {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			fmt.Println("Your name contains a vowel.")
			return
		}
	}
	fmt.Println("Your name does not contain a vowel.")
}
