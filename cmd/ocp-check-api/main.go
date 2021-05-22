package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enescakir/emoji"
	"github.com/ozoncp/ocp-check-api/core/api"
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

	for i := 0; i < 10000; i++ {
		func() {
			file, err := os.Open("go.mod")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
		}()
	}

	var check = &api.Check{}
	check.Init(1, 2, 3, 4, true)
	fmt.Println(check.String())
}
