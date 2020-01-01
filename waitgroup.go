package wait

import (
	"context"
	"fmt"
	"sync"
)

// Group extends sync.WaitGroup with a context aware wait method
// allowing early cancellation
type Group struct {
	sync.WaitGroup
}

// WaitWithContext blocks until the underlying WaitGroup counter is zero
// or the passed in context throws a cancellation signal
func (w *Group) WaitWithContext(ctx context.Context) error {
	done := make(chan struct{})

	go func() { w.Wait(); done <- struct{}{} }()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		fmt.Println("test")
		return ctx.Err()
	}
}
