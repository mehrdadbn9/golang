package main

import "fmt"

func init() {
	fmt.Println("init")
}

func main() {
	hello()
}

func hello() {
	fmt.Println("Hello, World!")
}
