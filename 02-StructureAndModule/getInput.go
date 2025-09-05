package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println("=== Method 1: Using bufio.Scanner ===")
    namesWithScanner()
    
    fmt.Println("\n=== Method 2: Using fmt.Scanln ===")
    namesWithScanln()
}

// Method 1: Using bufio.Scanner to handle full line input, including spaces
func namesWithScanner() {
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
    
    fmt.Printf("Method 1 captured: '%s'\n", name)
    
    // Check whether the name contains a vowel
    for _, v := range strings.ToLower(name) {
        if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
            fmt.Println("Your name contains a vowel.")
            return
        }
    }
    fmt.Println("Your name does not contain a vowel.")
}

// Method 2: Using fmt.Scanln to capture input
func namesWithScanln() {
    fmt.Println("Enter your name:")
    var name string
    fmt.Scanln(&name)
    
    fmt.Printf("Method 2 captured: '%s'\n", name)
    
    // Check whether name has a vowel
    for _, v := range strings.ToLower(name) {
        if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
            fmt.Println("Your name contains a vowel.")
            return
        }
    }
    fmt.Println("Your name does not contain a vowel.")
}