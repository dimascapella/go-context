package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContextWithDeadLine(t *testing.T) {
	fmt.Println("Goroutine = ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCount(ctx)
	for n := range destination {
		fmt.Println("Counter ", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Goroutine = ", runtime.NumGoroutine())
}
