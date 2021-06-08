package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ozoncp/ocp-check-api/internal/flusher"
	"github.com/ozoncp/ocp-check-api/internal/models"
	"github.com/ozoncp/ocp-check-api/internal/repo"
	"github.com/ozoncp/ocp-check-api/internal/saver"
)

func Greeting(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

// Чтение консольного ввода
func read(r io.Reader) <-chan string {
	consoleCh := make(chan string)
	go func() {
		defer close(consoleCh)
		scan := bufio.NewScanner(r)
		for scan.Scan() {
			s := scan.Text()
			consoleCh <- s
		}
	}()
	return consoleCh
}

func main() {
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	capacity := uint(10)
	chunkSize := 100
	repo := repo.NewCheckRepo()
	flusher := flusher.NewCheckFlusher(chunkSize, repo)

	saver := saver.NewSaver(capacity, time.Second*3, flusher)
	if saver != nil {
		saver.Init()
	}

	defer saver.Close()
	consoleCh := read(os.Stdin)

	go func() {
		fmt.Println("Enter new string to save models.Check or press Ctrl+C to exit...")
		for msg := range consoleCh {
			_ = saver.Save(models.Check{ID: uint64(len(msg)), Success: len(msg) > 2})
		}
	}()

	<-termChan // Ctrl+C pressed!
}
