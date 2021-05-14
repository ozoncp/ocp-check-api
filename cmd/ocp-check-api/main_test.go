package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"testing"
)

var binaryName = "test-ocp-check-api"

func TestMain(m *testing.M) {
	build := exec.Command("go", "build", "-o", binaryName)
	err := build.Run()
	if err != nil {
		fmt.Printf("could not make binary %v", err)
		os.Exit(1)
	}
	exitCode := m.Run()

	cleanUp := exec.Command("rm", "-f", binaryName)
	cleanUperr := cleanUp.Run()
	if cleanUperr != nil {
		fmt.Println("could not clean up", err)
	}

	os.Exit(exitCode)
}

func TestAbs(t *testing.T) {
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}
