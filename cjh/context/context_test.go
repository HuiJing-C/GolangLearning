package context

import (
	"context"
	"testing"
	"time"
)

func TestControl(t *testing.T) {
	cancel, _ := context.WithCancel(context.Background())
	deadline, _ := context.WithDeadline(cancel, time.Now().Add(time.Second*3))
	go func(ctx context.Context) {
		for {
			println(1)
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}(deadline)
	<-deadline.Done()
}
