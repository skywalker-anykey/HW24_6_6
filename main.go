package main

import "fmt"

func main() {
	var name string
	_, _ = fmt.Scanln(&name)
	fmt.Printf("Hello, %s!\n", name)
}
