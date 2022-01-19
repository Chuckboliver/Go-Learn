package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func slowProcess(ctx context.Context) {
	defer wg.Done()

	select {
	case <-time.After(20 * time.Second):
		fmt.Println("Task finish")
	case <-ctx.Done():
		fmt.Println("Canceled by context : ", ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	//timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	deadlineCtx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))

	defer cancel()

	fmt.Println("Start")
	wg.Add(1)
	go slowProcess(deadlineCtx)
	wg.Wait()
	fmt.Println("Finish")
}
