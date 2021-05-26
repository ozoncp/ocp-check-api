package main

import (
	"encoding/json"
	"flag"
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

	var checks = []api.Check{}
	if nc, _ := api.NewCheck(2, 3, 4, 5, false); nc != nil {
		checks = append(checks, *nc)
	}
	checks = append(checks, api.Check{ID: 3, SolutionID: 4, TestID: 5, RunnerID: 6, Success: true})

	str := `{"id": 4, "success": false}`
	jsonCheck := api.Check{}
	if err := json.Unmarshal([]byte(str), &jsonCheck); err == nil && jsonCheck.ID > 0 {
		checks = append(checks, jsonCheck)
	}

	batches := utils.SplitToBulks(checks, 10)
	fmt.Printf("Batches: %v\n", batches[0][0].String())
	fmt.Printf("Batch[0] len: %v\n", len(batches[0]))
}
