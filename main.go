package main

import "fmt"

func main() {
	fmt.Print("Hello world!")
	server := NewAPIServer(":3000")
	server.Run()
}
