package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/enescakir/emoji"
	"github.com/ozoncp/ocp-check-api/internal/models"
	"github.com/ozoncp/ocp-check-api/internal/utils"
)

func Greeting(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

func taskDeferInLoop() {
	for i := 0; i < 10; i++ {
		func() {
			file, err := os.Open("go.mod")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
		}()
	}
}
func main() {
	runDeferOpenPtr := flag.Bool("run-defer-task", false, "run deferred file closing in for loop")

	flag.Parse()

	if *runDeferOpenPtr {
		taskDeferInLoop()
	}

	fmt.Println(Greeting("Vladimir Cherdantsev"))
	fmt.Printf("It is a main %v\n", emoji.Package)

	{
		var checks = []models.Check{}
		if nc, _ := models.NewCheck(2, 3, 4, 5, false); nc != nil {
			checks = append(checks, *nc)
		}
		checks = append(checks, models.Check{ID: 3, SolutionID: 4, TestID: 5, RunnerID: 6, Success: true})

		str := `{"id": 4, "success": false}`
		jsonCheck := models.Check{}
		if err := json.Unmarshal([]byte(str), &jsonCheck); err == nil && jsonCheck.ID > 0 {
			checks = append(checks, jsonCheck)
		}

		batches := utils.SplitChecksToBulks(checks, 10)
		fmt.Printf("First check: %v\n", batches[0][0].String())
		fmt.Printf("Batch[0] len: %v\n", len(batches[0]))
	}

	{
		var tests = []models.Test{}
		if nc, _ := models.NewTest(7, 8, "run", "Hello world!"); nc != nil {
			tests = append(tests, *nc)
		}
		tests = append(tests, models.Test{ID: 3, TaskID: 4, Output: "wrong", Input: "try"})

		str := `{"id": 4, "taskID": 12}`
		jsonTest := models.Test{}
		if err := json.Unmarshal([]byte(str), &jsonTest); err == nil && jsonTest.ID > 0 {
			tests = append(tests, jsonTest)
		}

		batches := utils.SplitTestsToBulks(tests, 2)
		fmt.Printf("First test: %v\n", batches[0][0].String())
		fmt.Printf("Batch[0] len: %v\n", len(batches[0]))
	}
}
