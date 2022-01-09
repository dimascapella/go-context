package go_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithValue(t *testing.T) {
	contextParent := context.Background()

	contextChild1 := context.WithValue(contextParent, "Child1", "Child1")
	contextChild2 := context.WithValue(contextParent, "Child2", "Child2")

	contextChild11 := context.WithValue(contextChild1, "Child11", "Child11")
	contextChild12 := context.WithValue(contextChild1, "Child12", "Child12")

	contextChild21 := context.WithValue(contextChild2, "Child21", "Child21")

	fmt.Println(contextParent)
	fmt.Println(contextChild1)
	fmt.Println(contextChild2)
	fmt.Println(contextChild11)
	fmt.Println(contextChild12)
	fmt.Println(contextChild21)
	fmt.Println(contextChild21.Value("Child2"))
}
