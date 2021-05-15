package main

import (
	"fmt"

	"github.com/enescakir/emoji"
)

func Greeting(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

func main() {
	fmt.Println(Greeting("Vladimir Cherdantsev"))
	fmt.Printf("It is a main %v\n", emoji.Package)
}
