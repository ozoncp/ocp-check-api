package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enescakir/emoji"
)

func Greeting(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

func main() {
	fmt.Println(Greeting("Vladimir Cherdantsev"))
	fmt.Printf("It is a main %v\n", emoji.Package)

	for i := 0; i < 10000; i++ {
		func() {
			file, err := os.Open("go.mod")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
		}()
	}
}
