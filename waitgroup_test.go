package wait

import (
	"context"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	var wg Group

	wg.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
	}()

	if err := wg.WaitWithContext(ctx); err != context.DeadlineExceeded {
		t.Fatalf("err should be %v, got %v", context.DeadlineExceeded, err)
	}
}
