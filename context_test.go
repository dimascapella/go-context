package go_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	background := context.Background()
	todo := context.TODO()

	fmt.Println(background)
	fmt.Println(todo)
}
