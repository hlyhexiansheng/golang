package main

import "fmt"

func main() {
	greet("Jane", "Doe", "Doe")
}

func greet(fname, lname string,a string) {
	fmt.Println(fname, lname,a)
}
