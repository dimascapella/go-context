package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Goroutine = ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCount(ctx)
	for n := range destination {
		fmt.Println("Counter ", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Goroutine = ", runtime.NumGoroutine())
}

func CreateCount(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter += 1
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}
