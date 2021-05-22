package main

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/ozoncp/ocp-check-api/internal/utils"
)

func Greeting(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

type Any = utils.Any

func main() {
	fmt.Println(Greeting("Vladimir Cherdantsev"))
	fmt.Printf("It is a main %v\n", emoji.Package)

	// todo: move following lines into tests
	var slice = []string{"one", "two", "three"}
	const batchSize = 3
	var batch, err = utils.BatchSlice(slice, batchSize)
	if err != nil {
		panic(err)
	}
	fmt.Println(batch)

	var sourceMap = map[string]Any{"one": 2, "three": "four"}
	var transposedMap, _ = utils.TransposeMap(sourceMap)
	fmt.Println(transposedMap)

	var filtered = utils.Filter([]string{"one", "two", "three"}, []string{"four", "five", "two"})
	fmt.Println(filtered)
}
