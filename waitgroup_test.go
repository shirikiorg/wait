package wait

import (
	"context"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {

	testSuite := []struct {
		timeOutDuration, operationDuration time.Duration
		expectedErr                        error
	}{
		{
			timeOutDuration:   1 * time.Millisecond,
			operationDuration: 1 * time.Second,
			expectedErr:       context.DeadlineExceeded,
		},
		{
			timeOutDuration:   1 * time.Second,
			operationDuration: 1 * time.Millisecond,
			expectedErr:       nil,
		},
	}

	for _, test := range testSuite {
		ctx, cancel := context.WithTimeout(context.Background(), test.timeOutDuration)
		defer cancel()

		var wg Group

		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(test.operationDuration)
		}()

		if err := wg.WaitWithContext(ctx); err != test.expectedErr {
			t.Fatalf("err should be %v, got %v", test.expectedErr, err)
		}
	}
}
